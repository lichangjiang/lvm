FROM busybox
WORKDIR .

# TODO Change to your desired driver.
COPY plugin-deploy.sh /deploy.sh
COPY ./drivers/lvm /lvm

CMD /bin/sh /deploy.sh