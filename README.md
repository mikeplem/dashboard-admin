# Dashboard Admin

The idea of this repo is to be a redesign of the adminchrome code. I am thinking about running consul and confd on each TV. Consul would store the needed URL, and possible other information, while then using confd to control the browser reload or URL change.

I wonder if this would make it easier to have automatic TV changes occur. The adminchrome code works along with the remotechrome code to tell the TV what to do. My thinking is that the consul communication does the work of making sure the TV gets the latest information. This way the admin and remote side would do less work.

At the moment this code does nothing more than query the consul cluster to get information out of and into the cluster. This is only POC code at this time.

## References

### Dependencies

<https://github.com/hashicorp/consul>

<https://github.com/kelseyhightower/confd>

### Code being superseded

<https://github.com/mikeplem/adminchrome>

<https://github.com/mikeplem/remotechrome>


