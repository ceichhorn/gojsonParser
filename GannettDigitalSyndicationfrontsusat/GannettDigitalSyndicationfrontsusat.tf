variable "org" { default = "gannett" }

data "archive_file" "bundle" {
   type         = "zip"
   source_dir   = "${path.module}/proxy_files"
   output_path  = "${path.module}/proxy_files_bundle/apiproxy.zip"
}

# proxy
resource "apigee_api_proxy" "GannettDigitalSyndicationfrontsusat" {
   name  = "GannettDigitalSyndicationfrontsusat"
   bundle       = "${data.archive_file.bundle.output_path}"
   bundle_sha   = "${data.archive_file.bundle.output_sha}"
}

# staging proxy deployment
resource "apigee_api_proxy_deployment" "GannettDigitalSyndicationfrontsusat_staging_deployment" {
  proxy_name   = "${apigee_api_proxy.GannettDigitalSyndicationfrontsusat.name}"
  org          = "${var.org}"
  env          = "staging"
  revision     = "${apigee_api_proxy.GannettDigitalSyndicationfrontsusat.revision}"
}
