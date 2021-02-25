# Release Metrics Generator ![example workflow](https://github.com/smutil/release_metrics_generator/actions/workflows/build-actions.yml/badge.svg)

CLI to generate below release metrics based on tag and git commit log.
1. Release Name
2. Change Volume
3. Leadtime
4. Author
5. Release Date


Usage
-----
 step 1. download release_metrics_generator from <a href=https://github.com/smutil/release_metrics_generator/releases>releases</a>. 
 
 step 2. create [config.yml](https://github.com/smutil/release_metrics_generator/config.yml). If config.yaml and release_metrics_generator is not in same location, you can provide the config.yml path using --config
 
 step 3. execute the release_metrics_generator as shown below. 
 
 ```
 ./release_metrics_generator --config /path-to-config.yml
 ```
 step 4. ReleaseMetrics.html will be generated in same location.

 ![Alt text](docs/images/release_metrics.png?raw=true "Title")

