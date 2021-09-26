FROM alpine:3.14
COPY ./HW_1_Ancient_Crypto.csproj /HW/
COPY ./HW_1_Ancient_Crypto.sln /HW/
COPY ./Program.cs /HW/
COPY ./MyExtensions.cs /HW/

RUN apk update &&\
    apk add --upgrade wget &&\
    apk add xsel &&\
    apk add bash icu-libs krb5-libs libgcc libintl libssl1.1 libstdc++ zlib

RUN wget https://dot.net/v1/dotnet-install.sh &&\
    chmod a+rx ./dotnet-install.sh &&\
    ./dotnet-install.sh -c Current

ENV DOTNET_ROOT=/root/.dotnet

RUN /root/.dotnet/dotnet build /HW

ENTRYPOINT [ "/HW/bin/Debug/net5.0/HW_1_Ancient_Crypto" ]