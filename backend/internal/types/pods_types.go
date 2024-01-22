package types

import "time"

type PodListRes struct {
	Items []PodInfo `json:"items"`
}
type PodInfo struct {
	Metadata Metadata  `json:"metadata"`
	Spec     Spec      `json:"spec"`
	Status   PodStatus `json:"status"`
}

type Container struct {
	Name                     string `json:"name"`
	Image                    string `json:"image"`
	TerminationMessagePath   string `json:"terminationMessagePath"`
	TerminationMessagePolicy string `json:"terminationMessagePolicy"`
	ImagePullPolicy          string `json:"imagePullPolicy"`
}

type Spec struct {
	Containers                    []Container `json:"containers"`
	RestartPolicy                 string      `json:"restartPolicy"`
	TerminationGracePeriodSeconds int         `json:"terminationGracePeriodSeconds"`
	DnsPolicy                     string      `json:"dnsPolicy"`
	ServiceAccountName            string      `json:"serviceAccountName"`
	ServiceAccount                string      `json:"serviceAccount"`
	NodeName                      string      `json:"nodeName"`
	SchedulerName                 string      `json:"schedulerName"`
	Priority                      int         `json:"priority"`
	EnableServiceLinks            bool        `json:"enableServiceLinks"`
	PreemptionPolicy              string      `json:"preemptionPolicy"`
}

type ContainerStatus struct {
	Name         string `json:"name"`
	Ready        bool   `json:"ready"`
	RestartCount int    `json:"restartCount"`
	Image        string `json:"image"`
	ImageID      string `json:"imageID"`
	ContainerID  string `json:"containerID"`
	Started      bool   `json:"started"`
}

type PodStatus struct {
	Phase             string            `json:"phase"`
	HostIP            string            `json:"hostIP"`
	PodIP             string            `json:"podIP"`
	StartTime         string            `json:"startTime"`
	ContainerStatuses []ContainerStatus `json:"containerStatuses"`
}

type Metadata struct {
	Name              string    `json:"name"`
	GenerateName      string    `json:"generateName"`
	Namespace         string    `json:"namespace"`
	Uid               string    `json:"uid"`
	ResourceVersion   string    `json:"resourceVersion"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
}
