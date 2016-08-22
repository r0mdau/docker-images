FROM r0mdau/jessie:base
RUN apt-get update
RUN apt-get install -y x11vnc xvfb iceweasel
RUN mkdir ~/.vnc
RUN x11vnc -storepasswd password ~/.vnc/passwd
RUN bash -c 'echo "iceweasel" >> ~/.bashrc'
