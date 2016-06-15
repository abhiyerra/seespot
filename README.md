# SeeSpotStop

SeeSpotStop provides a utility for AWS Spot instance that handles the 
health check if used with an AWS ELB and also handles cleanup of the
instance when a termination notice is sent.

## Usage

```sh
seespotstop -health-port=8686 -health-path=/health -app-health="https://localhost:8080/health" -cleanup-task=/path/to/cleanup.sh
```

This should be run within an upstart or systemd unit file.

## Description

SeeSpotStop watches for a [termination notification](https://aws.amazon.com/blogs/aws/new-ec2-spot-instance-termination-notices/)
every 5 seconds and upon notification closes off the health 
check on the application and initiates some cleanup tasks.

This provides the following:

 - Make sure there are not new connections being 
   made while the instance is going down.
 - Take care of removing machine from ELBs via 
   the health check.
 - Take care of the Application health check.
 
## License

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
