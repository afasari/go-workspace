FROM envoyproxy/envoy:v1.25-latest
RUN apt-get update
COPY envoy.yaml /etc/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy.yaml