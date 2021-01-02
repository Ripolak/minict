# Minict

Minict is a minimal container runtime written in Go. It was made mainly for learning purposes and is intended to be as simple as possible. 

It's main intention is to be easily understandable to anyone who wishes to read it's code and see what goes into running containers in famous runtimes such as Containerd and full container-running platforms like Docker.

Minict runs OCI standard images and supports pulling images from existing registries. 

## Prerequisites
 * The `gpgme-devel` package must be installed on your system.
   * Run `sudo dnf install gpgme-devel` on RHEL-based distros (RHEL, CentOS, Fedora, etc.)
   * Run `sudo apt install libgpgme-dev` on Debian-based distros (Debian, Ubuntu, etc.)
 * Have `golang` and `git` installed.

## Building & Running
 * Clone this repository into your `$HOME/go/src` directory and run `cd $HOME/go/src/minict`
 * Run the `go get` command.
 * Run the `go build` command.
 * You should now have a `minict` executable in your directory. Run `chmod a+x minict` and it can now be used.
 * To use it from anywhere easily, move the `minict` executable to a directory that is in yout `PATH` variable. One such directory should be `/usr/bin`.

 ## Getting Started
  * Pulling an image:
  ```bash
  sudo minict pull --image ubuntu:20.04
  ```
  * Running a new container:
  ```bash
  sudo minict run --image ubuntu:20.04 --name ubuntu-ctr
  ```
  * Starting an existing container:
  ```bash
  sudo minict start --name ubuntu-ctr
  ```
  * Listing all images:
  ```bash
  >>> sudo minictl list-images

  [
	"alpine:latest",
	"ubuntu:20.04"
  ]
  ```
  * Listing all containers:
  ```bash
  >>> sudo minictl list-containers
  
  [
	"container1",
	"container2",
	"alpine-ctr",
	"ubuntu-ctr"
]
  ```
  * Removing an existing container:
  ```bash
  sudo minict rm --name ubuntu-ctr
  ```

## Running information
 * Container management is entirely based on the filesystem, and no DBs or separate inventories are used.
 * By default, minict will use the `/var/lib/minict` directory for everything that it stores on disk. This can be changed by setting the `MINICT_DIR` environment variable to a new location. For example:
 ```bash
 sudo MINICT_DIR=/new/location minict pull --image alpine:3.11
 ```

## Important Notes & Disclaimers:
 * Since this is intended to be small and minimal project, only basic OCI settings are supported. Currently, minict supports the following features:
   * Cmd / Entrypoint
   * Env variables
   * Mounts - Partial support, some mounts won't work and you might see mounting error when starting containers, but the basic mounts that are required for most containers are supported.
   * Hostname
* All other OCI settings are ignored.
* Currently, containers do not start with a networking namespace. I intend to add it in the near future.
* **Minict is in NO way meant to be used in production.** This is nothing more than a personal project and I take no responsibility for anything that using it might cause.
* While I tested this on my personal system and tried my best to cleanup everything once a container is removed, note that it does setup namespaces and mounts on the system, so **for completely risk-free usage I'd suggest using it on a VM**.
