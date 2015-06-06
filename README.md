# Kubernetes worker for ArchCI

## Introduction

Kubernetes worker is the driver to integrated with ArchCI and kubernetes.

With kubernetes worker, we can deploy CI tasks in kubernetes cluster without any effort. Notice that this worker is under development, we will update the repo if we make some progress.

## Usage

It's written in go so you can run it at most platforms.

```
cd ArchCI/kubernetes-workers
go run worker.go
```

Or simply download the binary.

```
go get github.com/ArchCI/kubernetes-workers
```

The worker will pull tasks from [archci](https://github.com/ArchCI/archci), please make sure it's running at first.
