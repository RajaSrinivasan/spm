# spm - Secure Package Manager

## Background

Secure Package Manager (spm) evolved out of a need to distribute software updates to connected embedded devices with no practical user interfaces. It is derived from an actual functioning system, retaining only the features that are of general applicability.

The goal was to be able to update the application software, distribute data files, and change configurations; being built utilizing the embedded linux platform yocto, eventually the goal incorporated the upgrade of the Operating System itself. This subsystem belongs in a larger context of a network distribution infrastructure and a resilient data transfer, installation and subsequent activation.

High level requirements for such a subsystem then is to:
- collect a set of files to be installed on a target system
- enumerate the set of instructions on how to utilize the set of files
- package them in an encrypted form - requiring a previously shared key to decrypt
- a mechanism to authenticate each of the files at the destination
- driving the execution of the installation steps

## Installation

### Systemwide Configuration

A configuration file is used to specify parameters applicable to all the packages. An example is:

```
#-----------------------------------------------------------------------
#     This is an example configuration file systemwide for spm.
#     Typical location $HOME/.spm.yaml
#     can be overridden with the --config flag
#-----------------------------------------------------------------------
#     Environment Variables
#     SPM_PKGPASSWORD   - the password for the encrytion of the spm file
#-----------------------------------------------------------------------
pubpkg: https://drive.google.com/
pubart: https://drive.aws.com/

package:
    format: tgz
    workarea: /tmp
```

Of the parameters above, the workarea is sometimes overwritten to point to different partitions - in the case of embedded systems with limited storage e.g. on sdcards.

## Configuration of individual packages

For each package that needs built, spm accepts a configuration file similar to:

```

package:
    name: ServicePack
    version: 1.2

contents:
    - from: /Volumes/Dev1/Ref/Books/acsac.pdf 
      to: /tmp/acsac.pdf 


preinstall:
    - go version
    - ls -l /tmp

postinstall:
    - ls /tmp
    - python --version
```
### Section: package
Mostly intended for documentation.
### Section: contents
The pair of from and to can be repeated any number of times in the contents section.
### Sections: preinstall and postinstall
Each entry is a shell command applicable in the target system - typically a linux system. Depending on the context during installation, the commands may have to provide complete paths.

## Usage

../bin/spm
Secure package manager helps prepare and distribute packages of applications 
and/or data.

Usage:
  spm [command]

Available Commands:
  build       Build a secure package
  help        Help about any command
  install     Install the package

Flags:
      --config string   config file (default is $HOME/.spm.yaml)
  -h, --help            help for spm
      --keep            keep workarea

Use "spm [command] --help" for more information about a command.

### Building a package

../bin/spm build --help
Create a secure package based on the configuration file provided.
Optionally push the artifact(s) to a distribution server. 
The first argument is the package spec file (ex spec.yaml)
Output package name is the second argument

Usage:
  spm build [flags]

Flags:
  -h, --help   help for build

Global Flags:
      --config string   config file (default is $HOME/.spm.yaml)
      --keep            keep workarea

### Installing a package

../bin/spm install --help
Install the package provided first verifying the integrity of the artifacts. Argument
        is the package (.spm)

Usage:
  spm install [flags]

Flags:
  -h, --help   help for install
      --show   extract and show the contents. do not install. Implies --keep

Global Flags:
      --config string   config file (default is $HOME/.spm.yaml)
      --keep            keep workarea

## Example Usage

### Package Configuration
In the following package, one file is packaged and distributed to the target system at a specific location. There are a few
shell commands specified to be executed before the file installation (Preinstall) and another set to execute after the file installations(Postinstall).

```
package:
    name: ServicePack
    version: 1.2

contents:
    - from: /Volumes/Dev1/Ref/Books/acsac.pdf 
      to: /tmp/acsac.pdf 

preinstall:
    - go version
    - ls -l /tmp

postinstall:
    - ls /tmp
    - python --version
```

### Build a package
```
../bin/spm build systest/sp.yaml systest/sp.spm
Home dir is /Users/rajasrinivasan
Using config file: /Users/rajasrinivasan/Prj/go/spm/example/.spm.yaml
Pkg publish url=https://drive.google.com/ Artifacts=https://drive.aws.com/
Pkg Password Thisisagoodpassword Workarea /tmp
2020/01/22 14:46:10 Building package for configuration file systest/sp.yaml
2020/01/22 14:46:10 Workarea created /tmp/spm400690855
2020/01/22 14:46:10 Created dir /tmp/spm400690855/tmp/spm400690855/contents and /tmp/spm400690855/artifacts
Loaded package File: systest/sp.yaml Name : ServicePack
2020/01/22 14:46:10 Copying file /Volumes/Dev1/Ref/Books/acsac.pdf to /tmp/spm400690855/contents
2020/01/22 14:46:10 Generating keys Private: /tmp/spm400690855/work/private.pem and Public: /tmp/spm400690855/contents/public.pem
2020/01/22 14:46:11 Created keypair /tmp/spm400690855/work/private.pem and /tmp/spm400690855/contents/public.pem
2020/01/22 14:46:11 Content file /tmp/spm400690855/contents/acsac.pdf
2020/01/22 14:46:11 Content file /tmp/spm400690855/contents/public.pem
2020/01/22 14:46:11 Files: [/tmp/spm400690855/contents/acsac.pdf /tmp/spm400690855/contents/public.pem]
2020/01/22 14:46:11 Signing using /tmp/spm400690855/work/private.pem of 2 files
2020/01/22 14:46:11 Loading private key /tmp/spm400690855/work/private.pem
2020/01/22 14:46:11 Signing /tmp/spm400690855/contents/acsac.pdf creating /tmp/spm400690855/contents/acsac.pdf.sig
2020/01/22 14:46:11 Datahash: 7f5b4e683df4120ddb5a2937259255e8cef209e16cd6fba948f964959e6c4eb5
Signature: 2d114a1617cc2e007fceb25691e788780a35b5fdc650e33f5dd30f27d9026e58a0b59cb1d26a4771993550f09a7005db33d5e0e8eedf710560ecafec1f4b3d306110afebc887fc6b69b6ac0ecd5ee3395468daf5c730e2f65fdffaecd89247fbd2f88309d7eeb79949b0e6afa26f333cec292bb32fe7a2ec374fc7d24a7146aeec6ce29ac380185502d2d71fcdc1e186e7b16ef9d69f773b2b5667d1e5721f88adba84372de42dd9b27d9a77b0fa8ca15af8fa63bab988e579f83809a2998893448fec59e3387627a842d407719c3d445b0645319ee9e87e1972d46ef052e2f79e633b570e948693ac3ca3d3b567585a88ca34d1b4cc16de4ccf65c7be560cfb
2020/01/22 14:46:11 Signing /tmp/spm400690855/contents/public.pem creating /tmp/spm400690855/contents/public.pem.sig
2020/01/22 14:46:11 Datahash: bf54b4450a7438a38e769cd14199d747fb6f6c58e5757c348a81d9f5f7c0d179
Signature: 247921b321eda649e1964f495faeb5de89318f6bffb428c4dcc402c6d4538fd0e39664199e445010103c45c9a0dd1d4a6476082e7ff931e0174de295acb71eecf0442a3a977a7cc76f347e81cab49feadf40de8e417fc059211f5fd6b3499c4f8d86faeaa8dbe605f40d0af4836272ad45aec7a67c3c935a1af02b3c2ff5b28f5d3e65f0b1de9fc86e9bb504b3014d4ce5936a948da224376d2912367c30d4c1ee5554930d4f32afae713fa8d85c334e9f3a99aaf5bdc4c98836173ec0cf2330ff8e3aa61190e40fd9327ed88afe029ecdb7a64002a720210891296588f01d900405050685095e4524fe212166ea99082400a9c9d645448dd5d5e2960e65dd02
2020/01/22 14:46:11 Unique Id created b27d94be-d489-4f2f-9dca-9b954564843f
2020/01/22 14:46:11 Saved manifest /tmp/spm400690855/contents/Packagefile
2020/01/22 14:46:11 Signing /tmp/spm400690855/contents/Packagefile with /tmp/spm400690855/work/private.pem to generate /tmp/spm400690855/contents/Packagefile.sig
2020/01/22 14:46:11 Loading private key /tmp/spm400690855/work/private.pem
2020/01/22 14:46:11 Signing /tmp/spm400690855/contents/Packagefile creating /tmp/spm400690855/contents/Packagefile.sig
2020/01/22 14:46:11 Datahash: f7f851d7031c00aa71d22028040023db6df6b4b2e29fe4664892f17b711c2f7b
Signature: 3e8e185cd031516b4bed7516ac75a91a3ae6924a08161445bb4d59f7186718b2b168be90a764e8d4e4bca4bf3c0b971db4a26981f53ae169dc789423df5119188af5b46cbaec89047200a33978323b56800a094e14dd4d255950b5f24a1e2d82d31298c5f7f4fde9240b99ea7539e34ea1d4705ae6a04c44c5bd64136b8c65f2f66383f3a0d3f406d8c45f73f31a3a058cd8c9f32f31f9f5eabb2e6ed3d33021c9f5a15a962921da1d948a65bd61390fc9cf57795d11cbc56b5383098c59f2313dee0d906229d4c6cfb2074a218660907802fb6978c7e3959ad6c8610a3cb42cce76b3ecaf992be94fe5c6cee08749718c46af9a11e38ab8109f03aaccf8d285
2020/01/22 14:46:11 Signed the Package file. Generated /tmp/spm400690855/contents/Packagefile.sig
2020/01/22 14:46:11 Created /tmp/spm400690855/work/sp.spm
2020/01/22 14:46:11 Adding Packagefile Size 390
2020/01/22 14:46:11 Adding Packagefile.sig Size 256
2020/01/22 14:46:11 Adding acsac.pdf Size 123519
2020/01/22 14:46:11 Adding acsac.pdf.sig Size 256
2020/01/22 14:46:11 Adding public.pem Size 418
2020/01/22 14:46:11 Adding public.pem.sig Size 256
2020/01/22 14:46:11 Created /tmp/spm400690855/work/sp.spm
2020/01/22 14:46:11 Encrypt from: /tmp/spm400690855/work/sp.spm to systest/sp.spm passphrase Thisisagoodpassword
2020/01/22 14:46:11 IV: d1aca762acc9cd49ca7d33b9558d9de8
2020/01/22 14:46:11 Created systest/sp.spm
2020/01/22 14:46:11 Removed /tmp/spm400690855
```

### Install the above package
```
../bin/spm install systest/sp.spm
Home dir is /Users/rajasrinivasan
Using config file: /Users/rajasrinivasan/Prj/go/spm/example/.spm.yaml
Pkg publish url=https://drive.google.com/ Artifacts=https://drive.aws.com/
Pkg Password Thisisagoodpassword Workarea /tmp
2020/01/22 14:47:17 Installing package /Users/rajasrinivasan/Prj/go/spm/systest/sp.spm
2020/01/22 14:47:17 Workarea created /tmp/spm839508907
2020/01/22 14:47:17 Created dir /tmp/spm839508907/tmp/spm839508907/contents and /tmp/spm839508907/artifacts
2020/01/22 14:47:17 Decrypt from: /Users/rajasrinivasan/Prj/go/spm/systest/sp.spm to /tmp/spm839508907/work/sp.spm passphrase Thisisagoodpassword
2020/01/22 14:47:17 32 bytes read for password
2020/01/22 14:47:17 16 bytes read for IV
2020/01/22 14:47:17 IV: d1aca762acc9cd49ca7d33b9558d9de8
2020/01/22 14:47:17 110232 bytes read
2020/01/22 14:47:17 110232 bytes written
2020/01/22 14:47:17 Decrypted /Users/rajasrinivasan/Prj/go/spm/systest/sp.spm to create /tmp/spm839508907/work/sp.spm
2020/01/22 14:47:17 Extracting Packagefile
2020/01/22 14:47:17 Extracting Packagefile.sig
2020/01/22 14:47:17 Extracting acsac.pdf
2020/01/22 14:47:17 Extracting acsac.pdf.sig
2020/01/22 14:47:17 Extracting public.pem
2020/01/22 14:47:17 Extracting public.pem.sig
Loaded package File: /tmp/spm839508907/contents/Packagefile Name : ServicePack
2020/01/22 14:47:17 Authenticating /tmp/spm839508907/contents/acsac.pdf signature /tmp/spm839508907/contents/acsac.pdf.sig publickey file /tmp/spm839508907/contents/public.pem
2020/01/22 14:47:17 Loading public key /tmp/spm839508907/contents/public.pem
2020/01/22 14:47:17 Public key file /tmp/spm839508907/contents/public.pem parsed
2020/01/22 14:47:17 Verified the signature /tmp/spm839508907/contents/acsac.pdf.sig of file /tmp/spm839508907/contents/acsac.pdf
2020/01/22 14:47:17 Authenticating /tmp/spm839508907/contents/Packagefile signature /tmp/spm839508907/contents/Packagefile.sig publickey file /tmp/spm839508907/contents/public.pem
2020/01/22 14:47:17 Loading public key /tmp/spm839508907/contents/public.pem
2020/01/22 14:47:17 Public key file /tmp/spm839508907/contents/public.pem parsed
2020/01/22 14:47:17 Verified the signature /tmp/spm839508907/contents/Packagefile.sig of file /tmp/spm839508907/contents/Packagefile
2020/01/22 14:47:17 Executing Preinstall steps
2020/01/22 14:47:17 go version go1.12.1 darwin/amd64

2020/01/22 14:47:17 lrwxr-xr-x@ 1 root  wheel  11 Feb  8  2019 /tmp -> private/tmp

2020/01/22 14:47:17 
2020/01/22 14:47:17 
2020/01/22 14:47:17 Executing Postinstall steps
2020/01/22 14:47:17 0E54B0DC-3D67-4903-99AE-F0D43B3655D2
0F0C0FE4-C809-42F5-A2B8-BFB1A097224E
2121E7F2-A1FC-4A1A-9BB1-7B1B3A919591
4ADB67E3-6507-4735-8226-E4B2AC35E3B6
612A3519-AB12-4F2F-9492-473A95074FA7
9C96170E-6DD2-4027-98D5-023DF5261272
9E4652A4-5EAE-4D2B-AAFD-10A030963615
AE0F797D-7E7D-4C33-A489-727249DB451B
BF9CEB3F-CFDC-46E4-AD51-D21CAE11203E
DC263519-51BB-46E1-BF23-08B90B157DAB
F21554BE-CD85-450C-938C-DCA11A67D796
FB48D565-A389-47AB-B70A-E14EDC97CF23
acsac.pdf
adobesmuoutp3XmNBJ
adobesmuoutp8HNNgv
adobesmuoutpVfsj1m
adobesmuoutpkqlY3T
com.apple.launchd.6lDvSEZYde
com.apple.launchd.JKHYdsQzV3
ext
powerlog
spm839508907

2020/01/22 14:47:17 Python 2.7.16

2020/01/22 14:47:17 Removed /tmp/spm839508907
```

## Design Choices

### Digital Signatures for individual files

spm generates a public and private key pair for every invocation. The private key is used to generate signature files for each of the content files. Then the private key file is discarded but the public key is saved in the package. Any tampering of any files then will be detected when the contents are authenticated with the signature files. Further details can be gleaned from sign.go.

### Container File

All the contents are packaged up in a compressed tar file ie .tgz. The detailed format can be gleaned from pack.go.

### Encryption of the container file

The current implementation uses the Output Feedback Mode [OFB](https://csrc.nist.gov/publications/detail/sp/800-38a/final). Further details can be gleaned from crypt.go.
