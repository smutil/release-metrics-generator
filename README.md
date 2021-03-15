# Release Metrics Generator 

![example workflow](https://github.com/smutil/release_metrics_generator/actions/workflows/build-actions.yml/badge.svg)![example workflow](https://github.com/smutil/release_metrics_generator/actions/workflows/release-actions.yml/badge.svg)

CLI to generate below release metrics based on tag and git commit log.
1. Application Name
2. Release Name (version)
3. Change Volume
4. Leadtime
5. Author
6. Release Date

This tool can also push the data to influxdb and you can create dashboard for change leadtime, change volume etc. Please refer below sample dashboard.

 ![Alt text](docs/images/lead_time.png?raw=true "Leadtime")

 ![Alt text](docs/images/change_volume.png?raw=true "ChangeVolume")

Prerequsites
------------
1. Production deployments must be done from Master/Main branch
2. Tag must be created immedialty after Production Deployment


Usage
-----
 step 1. download release_metrics_generator from <a href=https://github.com/smutil/release_metrics_generator/releases>releases</a>. 
 
 step 2. create [config.yml](https://github.com/smutil/release_metrics_generator/config.yml). If config.yaml and release_metrics_generator is not in same location, you can provide the config.yml path using --config
 
 step 3. execute the release_metrics_generator as shown below. 
 
 ```
 ./release_metrics_generator --config /path-to-config.yml
 ```
 step 4. ReleaseMetrics.html and ReleaseMetrics.json will be generated in same location.

 ![Alt text](docs/images/release_metrics.png?raw=true "Title")

