package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	//"os/exec"
	"regexp"
	"log"
	"strconv"
)

// Apis struct which contains
// an array of users
type Apis struct {
	Apis []Api `json:"apis"`
}

// Api struct which contains a name
// a type and a list of social links
type Api struct {

	Definition Definition `json:"api_definition"`

}

// Definitions struct which contains a
// list of api's
type Definition struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Slug  string `json:"slug"`
	Keyless string `json:"use_keyless"`
	Proxy Proxy   `json:"proxy"`
	Vdata Vdata	  `json:"version_data"`
  CacheV CacheV `json:"cache_options"`
}

// list of uri's
type Proxy struct {

	Listen string `json:"listen_path"`
	Target  string `json:"target_url"`
	Rlimit  bool   `json:"disable_rate_limit"`
	Slisten bool `json:strip_listen_path`
	ServiceD ServiceD `json:"service_discovery"`
}

// list of service
type ServiceD struct {
// CacheEnabled bool  `json:"enable_cache"`
CacheTime int `json:"cache_timeout"`
}

// list of cache
type CacheV struct {
  CacheEnable bool  `json:"enable_cache"`
	CacheTime int `json:"cache_timeout"`
 }

// list of version
type Vdata struct {
Nversion bool  `json:"not_versioned"`
Versions Versions `json:"versions"`
}
// list of links
type Versions struct {
  Default Default  `json:"Default"`
}
type Default struct {
	Dname string `json:"name"`
	ExtPaths bool `json:"use_extended_paths"`
  Extended Extended  `json:"extended_paths"`
}
type Extended struct {
	Rewrited string  `json:"url_rewrites"`
  Rewrite []Rewrites  `json:"url_rewrites"`
}
type Rewrites struct {
  Path string  `json:"path"`
	Match string  `json:"match_pattern"`
}

// fmt.Println("Api path: " + apis.Apis[i].Definition.Vdata.Versions.Default.Extended.Rewrites.Rwpath)
// fmt.Println("Api pattern: " + apis.Apis[i].Definition.Vdata.Versions.Default.Extended.Rewrites.Rwpatt)

func main() {


	// Open our jsonFile
	jsonFile, err := os.Open("outrewrite5.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println("Successfully Opened outrewrite5.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened json as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)


	// we initialize our Apis array
	var apis Apis


	// we unmarshal our byteArray which contains our
	// jsonFile's content which we defined above
	json.Unmarshal(byteValue, &apis)


	// Make a Regex to say we only want regular characters
	    reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	    if err != nil {
	        log.Fatal(err)
	    }


	// we iterate through every api within our api array and
	// print out the different variables
	// as just an example
	for i := 0; i < len(apis.Apis); i++ {
		//fmt.Println("Api Type: " + apis.Apis[i].Name)
		processedName := reg.ReplaceAllString(apis.Apis[i].Definition.Name, "")
		targetString := (apis.Apis[i].Definition.Proxy.Target)
		baseString := apis.Apis[i].Definition.Proxy.Listen
		//CacheYes := (apis.Apis[i].Definition.CacheV.CacheEnable)
		cacheTime := (apis.Apis[i].Definition.CacheV.CacheTime)

		// fmt.Println("A string of %s becomes %s \n", processedString)

	//	fmt.Println("Api Age: " + apis.Apis[i].Type)
		fmt.Println("-------------------  Api Name: " + processedName + "  ----------")
		// fmt.Println("Api Name: " + processedName)
		// fmt.Println("Api Id: " + apis.Apis[i].Definition.Id)
		fmt.Println("Api Slug: " + apis.Apis[i].Definition.Slug)
		// fmt.Println("Api Keyless: " + apis.Apis[i].Definition.Keyless)
		fmt.Println("Proxy Listen: " + apis.Apis[i].Definition.Proxy.Listen)
		fmt.Println("Proxy Target: " + apis.Apis[i].Definition.Proxy.Target)
		 fmt.Println("Proxy Strip: ", apis.Apis[i].Definition.Proxy.Slisten)
		fmt.Println("Api Versioned: ", apis.Apis[i].Definition.Vdata.Nversion)
		fmt.Println("Cache Cached: ", apis.Apis[i].Definition.CacheV.CacheEnable)
		fmt.Println("Service CacheTime: ", apis.Apis[i].Definition.CacheV.CacheTime)
		fmt.Println("Proxy Rate Limit Off: ", apis.Apis[i].Definition.Proxy.Rlimit)
		fmt.Println("Api Dname: " + apis.Apis[i].Definition.Vdata.Versions.Default.Dname)
		fmt.Println("Api ExtPaths: ",  apis.Apis[i].Definition.Vdata.Versions.Default.ExtPaths)
    fmt.Println("Api path: " + apis.Apis[i].Definition.Vdata.Versions.Default.Extended.Rewrited)
		// fmt.Println("Api path: " + apis.Apis[i].Definition.Vdata.Versions.Default.Extended.Rewrites.Path)
		// fmt.Println("Api pattern: " + apis.Apis[i].Definition.Vdata.Versions.Default.Extended.Rewrites.Match)

var Targ = fmt.Sprint(processedName + "/proxy_files/apiproxy/targets/default.xml")
var TargDir = fmt.Sprint(processedName + "/proxy_files/apiproxy/targets/")
var ProxyDir = fmt.Sprint(processedName + "/proxy_files/apiproxy/proxies/")
var PolicyDir = fmt.Sprint(processedName + "/proxy_files/apiproxy/policies/")
var ResourceDir = fmt.Sprint(processedName + "/proxy_files/apiproxy/resources/jsc")
var Proxy = fmt.Sprint(processedName + "/proxy_files/apiproxy/proxies/default.xml")
// var CacheYes = fmt.Sprint(apis.Apis[i].Definition.CacheV.CacheEnable)

   _ = os.MkdirAll(ProxyDir, 0755)
	 _ = os.MkdirAll(ResourceDir, 0755)
	 _ = os.MkdirAll(PolicyDir, 0755)
    _ = os.MkdirAll(TargDir, 0755)
		// var Targ = fmt.Sprint(processedName + "/proxy_files/apiproxy/targets/default.xml")
		///  create proxy file //

		input, err := ioutil.ReadFile("template/source.tf")
			 if err != nil {
							 log.Fatalln(err)
			 }

							// fmt.Println(input)

							 newContents := strings.Replace(string(input), "TEMPLATE-API", processedName, -1)

							// fmt.Println(newContents)

			var Dest =  fmt.Sprintf(processedName + "/" + processedName + ".tf")
			 err = ioutil.WriteFile(Dest, []byte(newContents), 0644)
			 if err != nil {
							 log.Fatalln(err)
			 }

  ////////  write out Proxies default file.

	proxies, err := ioutil.ReadFile("template/proxies.xml")
		 if err != nil {
						 log.Fatalln(err)
		 }

						 newProxy := strings.Replace(string(proxies), "TEMPLATE_BASE", baseString, -1)

		 err = ioutil.WriteFile(Proxy, []byte(newProxy), 0644)
		 if err != nil {
						 log.Fatalln(err)
		 }

		 ////////  write out Proxies template file.
		 var ProxyTemp = fmt.Sprint(processedName + "/proxy_files/apiproxy/" + processedName + "-api.xml")
	 	proxytemp, err := ioutil.ReadFile("template/apiproxy-template-api.xml")
	 		 if err != nil {
	 						 log.Fatalln(err)
	 		 }

               ProxyTemp2 := strings.Replace(string(proxytemp), "TEMPLATE-API", processedName, -1)
							 ProxyTemp3 := strings.Replace(string(ProxyTemp2), "LONGNAME", apis.Apis[i].Definition.Slug, -1)
	 						 newProxyTemp := strings.Replace(string(ProxyTemp3), "TEMPLATE_BASE", baseString, -1)


	 		 err = ioutil.WriteFile(ProxyTemp, []byte(newProxyTemp), 0644)
	 		 if err != nil {
	 						 log.Fatalln(err)
	 		 }

		 ////////  write out target file
		 target, err := ioutil.ReadFile("template/target.xml")
		 	 if err != nil {
		 					 log.Fatalln(err)
		 	 }

		 					 newTarget := strings.Replace(string(target), "TEMPLATE_TARGET", targetString, -1)

		 	 err = ioutil.WriteFile(Targ, []byte(newTarget), 0644)
		 	 if err != nil {
		 					 log.Fatalln(err)
		 	 }

			 ////////  write out url-rewrite xml
			 var Policy = fmt.Sprint(processedName + "/proxy_files/apiproxy/policies/url-rewrite.xml")
			 rewriteT, err := ioutil.ReadFile("template/url-rewrite.xml")
			 	 if err != nil {
			 					 log.Fatalln(err)
			 	 }

			 	 err = ioutil.WriteFile(Policy, []byte(rewriteT), 0644)
			 	 if err != nil {
			 					 log.Fatalln(err)
			 	 }

				 ////////  write out shared flow
				 var Flow = fmt.Sprint(processedName + "/proxy_files/apiproxy/policies/verify-shared-flow.xml")
				 flowT, err := ioutil.ReadFile("template/verify-shared-flow.xml")
				 	 if err != nil {
				 					 log.Fatalln(err)
				 	 }

				 	 err = ioutil.WriteFile(Flow, []byte(flowT), 0644)
				 	 if err != nil {
				 					 log.Fatalln(err)
				 	 }
					 ////////  write out respone-cache Policy

			 var Cache = fmt.Sprint(processedName + "/proxy_files/apiproxy/policies/response-caching.xml")
			 cacheT, err := ioutil.ReadFile("template/response-caching.xml")
				if err != nil {
								log.Fatalln(err)
				}

				err = ioutil.WriteFile(Cache, []byte(cacheT), 0644)
				if err != nil {
								log.Fatalln(err)
				}

		///   Write out thew rate limit policy.

		var Limit = fmt.Sprint(processedName + "/proxy_files/apiproxy/policies/rate-limit.xml")
		limitT, err := ioutil.ReadFile("template/rate-limit.xml")
		 if err != nil {
						 log.Fatalln(err)
		 }

		 err = ioutil.WriteFile(Limit, []byte(limitT), 0644)
		 if err != nil {
						 log.Fatalln(err)
		 }
		 
				////////  write out cache Proxies if it's cached
 			 	if apis.Apis[i].Definition.CacheV.CacheEnable {

					Stime := strconv.Itoa(cacheTime)
					var Cache = fmt.Sprint(processedName + "/proxy_files/apiproxy/policies/response-caching.xml")
		      cacheT, err := ioutil.ReadFile("template/response-caching.xml")
					fmt.Println(Stime, cacheTime)
		        if err != nil {
						 log.Fatalln(err)
		        }
				 newCache := strings.Replace(string(cacheT), "TIMEOUT", Stime, -1)
		       err = ioutil.WriteFile(Cache, []byte(newCache), 0644)
		       if err != nil {
						 log.Fatalln(err)
		         }
 				} else {
					 os.Remove(processedName + "/proxy_files/apiproxy/policies/response-caching.xml")
 						 fmt.Println("Cache Not Enabled: ")
			///  gotta pull in a default without caching
						 proxiesn, err := ioutil.ReadFile("template/proxiesnocache.xml")
					 		 if err != nil {
					 						 log.Fatalln(err)
					 		 }
					 						 newProxyn := strings.Replace(string(proxiesn), "TEMPLATE_BASE", baseString, -1)
					 		 err = ioutil.WriteFile(Proxy, []byte(newProxyn), 0644)
					 		 if err != nil {
					 						 log.Fatalln(err)
					 		 }
 				}
///   Done with chewckimg if caching is Enabled


	}
}
