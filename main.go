package main

import (
	"context"
	"fmt"
	"sort"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	// Load Kubernetes configuration from default location
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}

	// Create a Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Retrieve a list of all pods in the cluster
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Create a slice of restart times for each pod
	type podRestart struct {
		namespace    string
		podName      string
		restartTimes []time.Time
	}

	podRestarts := make([]podRestart, 0, len(pods.Items))

	for _, pod := range pods.Items {
		var restartTimes []time.Time

		for _, restart := range pod.Status.ContainerStatuses {
			if restart.State.Terminated != nil {
				restartTimes = append(restartTimes, restart.State.Terminated.FinishedAt.Time)
			}
		}

		if len(restartTimes) > 0 {
			sort.Slice(restartTimes, func(i, j int) bool {
				return restartTimes[i].After(restartTimes[j])
			})
			podRestarts = append(podRestarts, podRestart{
				namespace:    pod.ObjectMeta.Namespace,
				podName:      pod.ObjectMeta.Name,
				restartTimes: restartTimes,
			})
		}
	}

	// Sort the podRestarts slice by the most recent restart time
	sort.Slice(podRestarts, func(i, j int) bool {
		return podRestarts[i].restartTimes[0].After(podRestarts[j].restartTimes[0])
	})

	// Determine how many pods to show
	numPodsToShow := 10
	if numPodsToShow > len(podRestarts) {
		numPodsToShow = len(podRestarts)
	}

	// Print the list of pods with the most recent restart times as a table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pod", "Namespace", "Most Recent Restart"})

	for i := 0; i < numPodsToShow; i++ {
		podRestart := podRestarts[i]
		row := []string{podRestart.podName, podRestart.namespace, podRestart.restartTimes[0].Format("Jan 2 15:04:05 2006")}
		table.Append(row)
	}

	fmt.Printf("The %d pods with the most recent restart times are:\n", numPodsToShow)
	table.Render()
}
