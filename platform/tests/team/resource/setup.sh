#!/bin/bash

amp org switch org
amp stack up -c examples/stacks/pinger/pinger.yml
amp stack up -c examples/stacks/pinger/pinger.yml pi
