package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
)

type apiResource struct {
	r  metav1.APIResource
	gv schema.GroupVersion
}

type resourceNameLookup map[string][]apiResource

type resourceMap struct {
	list []apiResource
	m    resourceNameLookup
}

func getAPIs(client discovery.DiscoveryInterface) (*resourceMap, error) {
	start := time.Now()
	resList, err := client.ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	rm := &resourceMap{
		m: make(resourceNameLookup),
	}

	
}

func main() {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	configFlags := genericclioptions.NewConfigFlags(true)

	// use the current context in kubeconfig
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	start := time.Now()
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, ns := range namespaces.Items {
		pods, err := clientset.CoreV1().Pods(ns.GetName()).List(context.TODO(), metav1.ListOptions{})

		deployements, err := clientset.AppsV1().Deployments(ns.GetName()).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		daemonsets, err := clientset.AppsV1().DaemonSets(ns.GetName()).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Namespace: %s\n", ns.GetName())
		fmt.Printf("Total pod count: %d\n", len(pods.Items))

		for _, p := range pods.Items {
			fmt.Println(p.GetName(), "==>", p.Status.Phase)
		}

		fmt.Printf("Total deployment count: %d\n", len(deployements.Items))
		fmt.Printf("Total daemonset count: %d\n", len(daemonsets.Items))
		fmt.Println("====================================")
	}

	// list pod name
	// for _, p := range pods.Items {
	// 	fmt.Println(p.GetNamespace(), "==>", p.GetName(), "==>", p.Status.Phase)
	// 	for _, c := range p.Spec.Containers {
	// 		fmt.Println("cpu: ", c.Resources.Limits.Cpu().ToDec().AsDec(), "memory: ", c.Resources.Limits.Memory())
	// 	}
	// }
	color.Green("Time elapsed: %v", time.Since(start))
}
