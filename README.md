# SeeSpotStop

SeeSpotStop provides a utility for AWS Spot instance that handles the
health check if used with an AWS ELB and also handles cleanup of the
instance when a termination notice is sent. When the termination
notice is sent there is a 2 minute period before termination this
handles the cleanup.

## Usage

```sh
seespotstop -health-port=8686 -health-path=/health -app-health="https://localhost:8080/health" -cleanup-task=/path/to/cleanup.sh
```

```sh
$ seespotstop -help
Usage of ./seespotstop:
  -app-health string
        Application health check (default "http://127.0.0.1:8080/health")
  -cleanup-task string
        Script to run upon termination
  -health-path string
        Default health path the Load Balancer hits (default "/health")
  -health-port string
        Default health port to use with Load Balancers (default ":8686")
```

This should be run within an upstart or systemd unit file.

## Description

SeeSpotStop watches for a
[termination notification](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/)
every 5 seconds and upon notification closes off the health check on
the application and initiates the cleanup tasks. If no termination
notice is found then it continues to run.

This provides the following:

 - Make sure there are not new connections being made while the
   instance is going down.
 - Take care of removing machine from ELBs via the health check.
 - Take care of the Application health check.

## License

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
