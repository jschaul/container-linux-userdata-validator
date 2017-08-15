FROM scratch
MAINTAINER CoreOS

WORKDIR /opt/validate
ENTRYPOINT ["bin/validate"]

ADD bin/validate /opt/validate/bin/validate
