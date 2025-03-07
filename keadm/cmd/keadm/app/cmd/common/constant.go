/*
Copyright 2019 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

const (
	// KubeEdgeVersion sets the version of KubeEdge to be used
	KubeEdgeVersion = "kubeedge-version"

	// ImageRepository sets the image repository to pull images
	ImageRepository = "image-repository"

	// KubeConfig sets the path of kubeconfig
	KubeConfig = "kube-config"

	// Master sets the address of K8s master
	Master = "master"

	// CloudCoreIPPort sets the IP and port of KubeEdge cloud component
	CloudCoreIPPort = "cloudcore-ipport"

	// EdgeNodeName is KubeEdge node unique identification string
	EdgeNodeName = "edgenode-name"

	// RemoteRuntimeEndpoint is KubeEdge remote-runtime-endpoint string
	RemoteRuntimeEndpoint = "remote-runtime-endpoint"

	// CertPath sets the path of the certificates generated by the KubeEdge Cloud component
	CertPath = "certPath"

	// DefaultK8SMinimumVersion is the minimum version of K8S
	DefaultK8SMinimumVersion = 11

	// RuntimeType is default runtime type
	RuntimeType = "runtimetype"

	// DefaultKubeEdgeVersion is the default KubeEdge version, it must have no prefix 'v'
	DefaultKubeEdgeVersion = "1.14.0"

	// Token sets the token used when edge applying for the certificate
	Token = "token"

	// CertPort is the port where to apply for the edge certificate
	CertPort = "certport"

	AdvertiseAddress = "advertise-address"

	TokenSecretName = "tokensecret"

	TokenDataName = "tokendata"

	DomainName = "domainname"

	Labels = "labels"

	// CGroupDriver is type of edgecore Cgroup
	CGroupDriver = "cgroupdriver"

	// TarballPath sets the temp directory path for KubeEdge tarball, if not exist, download it
	// eg.  "/tmp/kubeedge" or "/etc/kubeedge" by default
	TarballPath = "tarballpath"

	StrCheck    = "check"
	StrDiagnose = "diagnose"

	// Init-beta flags combined below:

	// Allow appending manifests paths of manifests to keadm, separated by commas
	Manifests = "manifests"

	// Allow appending manifests paths of manifests to keadm, separated by commas, another supported flag
	Files = "files"

	// Dry-run flag
	DryRun = "dry-run"

	// Forced install
	Force = "force"

	// Skip CRDs
	SkipCRDs = "skip-crds"

	// External Helm Root
	ExternalHelmRoot = "external-helm-root"

	// Helm action
	HelmInstallAction  = "install"
	HelmManifestAction = "manifest"

	CmdGetDNSIP         = "cat /etc/resolv.conf | grep nameserver | grep -v -E ':|#' | awk '{print $2}' | head -n1"
	CmdGetStatusDocker  = "systemctl status docker |grep Active | awk '{print $2}'"
	CmdPing             = "ping %s -w %d |grep 'packets transmitted' |awk '{print $6}'"
	CmdGetMaxProcessNum = "sysctl kernel.pid_max|awk '{print $3}'"
	CmdGetProcessNum    = "ps -A|wc -l"

	EdgecoreConfig = "config"

	EdgeCoreServer = "127.0.0.1:10350"

	// EdgecoreConfigPath is default edgecore config path
	EdgecoreConfigPath = "/etc/kubeedge/config/edgecore.yaml"

	// CmdCopyFile is the cmd to copy file
	CmdCopyFile = "cp -r %s %s/"

	/*system info*/
	CmdDiskInfo    = "df -h > %s/disk"
	CmdArchInfo    = "arch > %s/arch"
	CmdProcessInfo = "ps -axu > %s/process"
	CmdDateInfo    = "date > %s/date"
	CmdUptimeInfo  = "uptime > %s/uptime"
	CmdHistorynfo  = "history -a && cat ~/.bash_history  > %s/history"
	CmdNetworkInfo = "netstat -pan > %s/network"

	PathCpuinfo   = "/proc/cpuinfo"
	PathMemory    = "/proc/meminfo"
	PathHosts     = "/etc/hosts"
	PathDNSResolv = "/etc/resolv.conf"

	/*edgecore info*/
	PathEdgecoreService = "/lib/systemd/system/edgecore.service"
	CmdEdgecoreVersion  = "edgecore  --version > %s/version"

	/*runtime info*/
	CmdDockerVersion    = "docker version > %s/version"
	CmdContainerInfo    = "docker ps -a > %s/containerInfo"
	CmdContainerLogInfo = "journalctl -u docker  > %s/log"
	CmdDockerInfo       = "docker info > %s/info"
	CmdDockerImageInfo  = "docker images > %s/images"
	PathDockerService   = "/lib/systemd/system/docker.service"

	DescAll     = "Check all item"
	DescArch    = "Check whether the architecture can work"
	DescCPU     = "Check node CPU requirements"
	DescMemory  = "Check node memory requirements"
	Descdisk    = "Check node disk requirements"
	DescDNS     = "Check whether DNS can work"
	DescRuntime = "Check whether runtime can work"
	DescNetwork = "Check whether the network is normal"
	DescPID     = "Check node PID requirements"

	/**Diagnose**/
	ArgDiagnoseNode  = "node"
	DescDiagnoseNode = "Diagnose edge node"

	ArgDiagnosePod  = "pod"
	DescDiagnosePod = "Diagnose pod"

	ArgDiagnoseInstall  = "install"
	DescDiagnoseInstall = "Diagnose install"
	/****/

	ArgCheckAll     = "all"
	ArgCheckArch    = "arch"
	ArgCheckCPU     = "cpu"
	ArgCheckMemory  = "mem"
	ArgCheckDisk    = "disk"
	ArgCheckDNS     = "dns"
	ArgCheckRuntime = "runtime"
	ArgCheckNetwork = "network"
	ArgCheckPID     = "pid"

	KB = 1024
	MB = KB * 1024
	GB = MB * 1024

	AllowedValueCPU     = 1
	AllowedValueMemory  = 256 * MB
	AllowedValueDisk    = GB
	AllowedValuePIDRate = 0.05

	AllowedCurrentValueCPURate  = 0.9
	AllowedCurrentValueMemRate  = 0.9
	AllowedCurrentValueDiskRate = 0.9

	AllowedCurrentValueMem  = 128 * MB
	AllowedCurrentValueDisk = 512 * MB
)

var (
	CheckObjectMap = []CheckObject{
		{
			Use:  ArgCheckAll,
			Desc: DescAll,
		},
		{
			Use:  ArgCheckCPU,
			Desc: DescCPU,
		},
		{
			Use:  ArgCheckMemory,
			Desc: DescMemory,
		},
		{
			Use:  ArgCheckDisk,
			Desc: Descdisk,
		},
		{
			Use:  ArgCheckDNS,
			Desc: DescDNS,
		},
		{
			Use:  ArgCheckRuntime,
			Desc: DescRuntime,
		},
		{
			Use:  ArgCheckNetwork,
			Desc: DescNetwork,
		},
		{
			Use:  ArgCheckPID,
			Desc: DescPID,
		},
	}

	DiagnoseObjectMap = []DiagnoseObject{
		{
			Use:  ArgDiagnoseNode,
			Desc: DescDiagnoseNode,
		},
		{
			Use:  ArgDiagnosePod,
			Desc: DescDiagnosePod,
		},
		{
			Use:  ArgDiagnoseInstall,
			Desc: DescDiagnoseInstall,
		},
	}

	// DefaultKubeConfig is the default path of kubeconfig
	// make it an var so it can be changed to adapt to windows(In rare cases, user name is Administrator)
	DefaultKubeConfig = "/root/.kube/config"
)
