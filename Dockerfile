FROM alpine

RUN apk add smartmontools
RUN apk add bash
RUN mkdir -p /opt
COPY opt/smart-mon-script.sh /opt/smart-mon-script.sh
RUN chmod +x /opt/smart-mon-script.sh
COPY kube-smart-mon /