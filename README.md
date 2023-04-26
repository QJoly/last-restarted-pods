# Last restarted pods

This Go script uses the Kubernetes Go client to check which pods in a cluster have been restarted most recently. It retrieves a list of all pods in the cluster, and then looks for containers in each pod that have been restarted. For each pod with at least one restart, it determines the most recent restart time. Finally, it outputs a list of the pods with the most recent restart times, sorted by restart time in descending order.

## Requirements

- Go 1.16 or later
- A running Kubernetes cluster with authentication credentials set up (e.g., `~/.kube/config` file)

## Installation

### Option 1: Download a pre-compiled binary

1. Go to the [releases page](https://github.com/qjoly/last-restarted-pods/releases) on GitHub.
2. Download the binary for your operating system.
3. Make the binary executable: `chmod +x last-restarted-pods`

### Option 2: Build from source

1. Clone the repository: `git clone https://github.com/qjoly/last-restarted-pods.git`
2. Change into the project directory: `cd last-restarted-pods`
3. Build the Go binary: `go build -o last-restarted-pods`
4. Make the binary executable: `chmod +x last-restarted-pods`

## Usage

To run the script, simply execute the binary with no arguments:

```bash
./last-restarted-pods
```

The script will retrieve a list of all pods in the cluster and output a table of the 10 pods with the most recent restart times, sorted by restart time in descending order. The output will include the pod name, namespace, and most recent restart time.

Example of output:

```bash
The 5 pods with the most recent restart times are:
+-----------------------------------+-----------------+----------------------+
|                POD                |    NAMESPACE    | MOST RECENT RESTART  |
+-----------------------------------+-----------------+----------------------+
| crash-pod-2                       | monitoring      | Apr 26 10:15:33 2023 |
| c-b0lwyx-28041120--1-lz7zc        | longhorn-system | Apr 26 02:00:09 2023 |
| c-6d4ixb-28041000--1-vzg4g        | longhorn-system | Apr 26 00:01:17 2023 |
| helm-install-traefik--1-xmgpd     | kube-system     | Aug 8 16:42:53 2022  |
| helm-install-traefik-crd--1-kvvt5 | kube-system     | Aug 8 16:42:44 2022  |
+-----------------------------------+-----------------+----------------------+
```

## Contributing

Contributions are welcome! If you find a bug or have an idea for a new feature, please open an issue or submit a pull request.

## License

This code is licensed under the MIT License. See the LICENSE file for details.
