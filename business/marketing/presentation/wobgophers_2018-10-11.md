---
layout: presentation_v1.0.0
title: "Go in der Google Cloud"
excerpt: "2018-10-11 @ wobgophers"
author: "Christian Prior-Mamulyan"
dummycontent: false
theme: solarized
---

<!-- meta                                               -->
<section>

<section>
<h1>&lt;meta&gt;</h1>
</section>

<section style="text-align: center;">
<h2>Slides available</h2>
<p><a href="{{ "/business/marketing/presentation/wobgophers_2018-10-11" | prepend: site.baseurl }}" >
https://cprior.github.io/go-gae-intro/</a></p>
</section>

<section>
<h1>&lt;/meta&gt;</h1>
</section>

</section>
<!-- /meta                                              -->

<!-- Google Cloud                                       -->
<section>

<section>
<h1>&lt;cloud id="google"&gt;</h1>
</section>

<section>
<h2>Goocle Cloud Platform</h2>

<div style="text-align: left;">

<p>The GCP offers several <a href="https://cloud.google.com/products/">products and services</a>:</p>

<ul>
<li>Cloud Computing: Compute, Storage, Databases, Networking, …</li>
<li>Analytics: Data analytics, AI & machine learning, …</li>
<li>Identity and security: Cloud identity, Security, …</li>
<li>Collaboration and Productivity: G Suite, …</li>
<li>Google Maps Platform: sic! …</li>
</ul>

</div>

</section>

<section>
<h2>Go in the cloud</h2>

<div>
<img src="{{ "/technology/logical/images/serverless-FaaS.png" | prepend: site.baseurl }}" alt="serverless cloud" width="" style="float: right; border: 0px; margin-right: 10%">
<p style="text-align: left;">To run Go in the cloud</p>
<ul style="width: 40%;">
<li>PAAS: Google App Engine GAE, IBM Cloud Foundry Apps</li>
<li>FAAS: Amazon Lambda, Google Cloud Functions</li>
</ul>
</div>

<p>GCP / Compute: <a href="https://cloud.google.com/appengine/">Google App Engine GAE</a></p>

</section>


<section>
<h2>Go – amongst other things</h2>

<div style="float: right; width:96px; margin-right: 10%">
<p style="text-align: left; font-size: 33%;"><img src="{{ "/technology/physical/images/gophercolor_byReneeFrench_64x64.png" | prepend: site.baseurl }}" alt="serverless cloud" width="" style="border: 0px;" /><br />
The Go gopher was designed by Renee French.</p>
</div>
<div style="">
<ul style="width: 60%;">
<li>Java, PHP, Node.js, Python, C#, .Net, Ruby and Go – or Docker</li>
<li>Application Versioning</li>
<li>Traffic splitting</li>
<li>Monitoring, Logging, Debugging</li>
</ul>
<p><em>Use cases e.g. web apps or mobile backends</em></p>
</div>

</section>

<section>
<h2>Environments <a href="https://cloud.google.com/appengine/docs/the-appengine-environments">🔗</a></h2>

<div style="width: 44%; float: right;">
<h3 style="text-decoration: underline;">Standard</h3>
<ul style="">
<li>Pricing Based on usage of vCPU, memory, and persistent disks</li>
<li>sandbox with runtime of a supported language</li>
<li>Go 1.6, 1.8, and 1.9 *</li>
</ul>
</div>
<div style="width: 44%; float: left;">
<h3 style="text-decoration: underline;">Flexible</h3>
<ul style="">
<li>Pricing based on instance hours; for free or at low cost</li>
<li>Docker container within Google Compute Engine VM</li>
<li>Go 1.8, 1.9, and 1.10 *</li>
</ul>
</div>

<div style="width: 100%; float: left;"><p>* as per 2018-10-10</p></div>



</section>

<section>
<h1>&lt;/cloud&gt;</h1>
</section>
</section>

<!-- /Google Cloud                                      -->

<!-- Hello World                                        -->
<section>

<section>
<h1>&lt;hello name="world"&gt;</h1>
</section>

<section data-markdown>
## Hello World ##

```Go
package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	appengine.Main()
}
```

</section>

<section>
<h2>Hello $world</h2>
<pre class="stretch"><code data-trim data-noescape Go>
package main
import ...

func helloHandler(w http.ResponseWriter, r *http.Request) {

	name := strings.Split(strings.TrimPrefix(
				r.URL.Path, "/hello"),
				"/")[1]

	if name != "" {
		fmt.Fprintf(w, "Hello %v!", name)
	} else {
		fmt.Fprintln(w, "Hello world!")
	}
}
func main() {
	http.HandleFunc("/hello/", helloHandler)
	appengine.Main()
}</code></pre>
</section>

<section>
<h1>&lt;/hello&gt;</h1>
</section>

</section>
<!-- /Hello World                                       -->

<!-- inspection                                         -->
<section>

<section>
<h1>&lt;inspection&gt;</h1>
</section>



<section>

<h2>References</h2>

<div class="fragment current-visible"><p><img style="border: 0px; width: 640px; height: 480px; background: url({{ '/application/logical/images/AppEngine-1-Create.png' | prepend: site.baseurl  }}) 0 0px;" src="{{ '/business/marketing/presentation/images/transparent.png' }}"/><br />
Create a project</p></div>

<div class="fragment current-visible"><p><img style="border: 0px; width: 640px; height: 480px; background: url({{ '/application/logical/images/AppEngine-2-NewProject.png' | prepend: site.baseurl  }}) 0 0px;" src="{{ '/business/marketing/presentation/images/transparent.png' }}"/><br />
Naming the project</p></div>

<div class="fragment current-visible"><p><img style="border: 0px; width: 640px; height: 480px; background: url({{ '/application/logical/images/AppEngine-3-SelectRegion.png' | prepend: site.baseurl  }}) 0 0px;" src="{{ '/business/marketing/presentation/images/transparent.png' | prepend: site.baseurl  }}" /><br />
Select the region where to deploy to</p></div>

<div class="fragment current-visible"><p><img style="border: 0px; width: 640px; height: 480px; background: url({{ '/application/logical/images/AppEngine-4-Preparing.png' | prepend: site.baseurl  }}) 0 0px;" src="{{ '/business/marketing/presentation/images/transparent.png' | prepend: site.baseurl  }}" /><br />
Wait for the project being prepared</p></div>

<div class="fragment current-visible"><p><img style="border: 0px; width: 640px; height: 480px; background: url({{ '/application/logical/images/AppEngine-5-YourFirstApp.png' | prepend: site.baseurl  }}) 0 0px;" src="{{ '/business/marketing/presentation/images/transparent.png' | prepend: site.baseurl  }}" /><br />
The beginning of the GAE app</p></div>
</section>

<section>
<h2>The SDK: installation <a href="https://cloud.google.com/sdk/">🔗</a></h2>

<div style="">
<p style="">Tools for the Google Cloud Platform</p>
<ul style="">
<li>Manage products and services</li>
<li>deploy code</li>
<li>Run local emulators</li>
</ul>

</div>

</section>

<section>
<h2>The SDK: components</h2>
<pre class="stretch"><code data-trim data-noescape bash>
gcloud components install app-engine-go

gcloud components list --only-local-state 

Your current Cloud SDK version is: 220.0.0
┌─────────────────────────────────────────────────────────────┐
│                           Components                        │
├──────────────────────────────────┬──────────────────────────┤
│                      Name        │            ID            │
├──────────────────────────────────┼──────────────────────────┤
│ Cloud SDK Core Libraries         │ core                     │
│ gcloud app Python Extensions     │ app-engine-python        │
│ gcloud app Python Extensions     │ app-engine-python-extras │
│ gcloud Beta Commands             │ beta                     │
│ Cloud Storage Command Line Tool  │ gsutil                   │
│ BigQuery Command Line Tool       │ bq                       │
│ App Engine Go Extensions         │ app-engine-go            │
│ Cloud Datastore Emulator         │ cloud-datastore-emulator │
└──────────────────────────────────┴──────────────────────────┘
</code></pre>

</section>

<section>
<h2>The SDK: update</h2>
<pre class="stretch"><code data-trim data-noescape Bash>
gcloud components update

Your current Cloud SDK version is: 211.0.0
You will be upgraded to version: 220.0.0

┌────────────────────────────────────────────────────────────────┐
│               These components will be updated.                │
├───────────────────────────────────────┬────────────┬───────────┤
│                         Name          │  Version   │    Size   │
├───────────────────────────────────────┼────────────┼───────────┤
│ App Engine Go Extensions              │     1.9.68 │ 153.3 MiB │
│ BigQuery Command Line Tool            │     2.0.34 │   < 1 MiB │
│ BigQuery Command Line Tool            │     2.0.34 │   < 1 MiB │
│ Cloud Datastore Emulator              │      2.0.2 │  17.7 MiB │
│ Cloud SDK Core Libraries              │ 2018.10.08 │   8.6 MiB │
│ Cloud SDK Core Libraries              │ 2018.09.24 │   < 1 MiB │
│ Cloud Storage Command Line Tool       │       4.34 │   3.5 MiB │
│ Cloud Storage Command Line Tool       │       4.34 │   < 1 MiB │
│ gcloud app Python Extensions          │     1.9.77 │   6.2 MiB │
│ gcloud app Python Extensions          │     1.9.74 │  28.5 MiB │
│ gcloud cli dependencies               │ 2018.09.28 │   2.4 MiB │
└───────────────────────────────────────┴────────────┴───────────┘
</code></pre>

</section>

<section>
<h2>Live Demo</h2>

<div style="text-align: left;">

<p>Let me try to cover the following aspects of a Go app on Google Apps Engine GAE:</p>

<ul>
<li>deployment</li>
<li>task queues</li>
<li>scheduled tasks</li>
<li>logs</li>
<li>endpoints</li>
<li>storage *</li>
<li>email *</li>
</ul>

<p>* if time permits</p>

</div>

</section>


<section>
<h1>&lt;/inspection&gt;</h1>
</section>

</section>
<!-- /inspection                                       -->

<!-- end                                               -->
<section>

<section>
<h1>&lt;end /&gt;</h1>
</section>

</section>
<!-- /end                                               -->

