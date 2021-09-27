FROM mcr.microsoft.com/dotnet/runtime:5.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:5.0 AS build
WORKDIR /src
COPY ["Diffie-Hellman/Diffie-Hellman.csproj", "Diffie-Hellman/"]
RUN dotnet restore "Diffie-Hellman/Diffie-Hellman.csproj"
COPY . .
WORKDIR "/src/Diffie-Hellman"
RUN dotnet build "Diffie-Hellman.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Diffie-Hellman.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Diffie-Hellman.dll"]
