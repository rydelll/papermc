# What file to download
- GET https://api.papermc.io/v2/projects/velocity
- Parse as JSON
- veloctiy_ver = response.versions[-1]
- GET https://api.papermc.io/v2/projects/velocity/versions/${velocity_ver}/builds
- Parse as JSON
- build_ver = response.builds[-1].build
- jar = response.builds[-1].downloads.application.name
- jar_sha = response.builds[-1].downloads.application.sha256
- GET https://api.papermc.io/v2/projects/velocity/versions/${velocity_ver}/builds/${build_ver}/downloads/${jar}/


## Useful links
https://api.papermc.io/v2/projects/velocity
https://api.papermc.io/v2/projects/velocity/versions/${velocity_ver}/builds
https://api.papermc.io/v2/projects/velocity/versions/${velocity_ver}/builds/${build_ver}/
https://api.papermc.io/v2/projects/velocity/versions/${velocity_ver}/builds/${build_ver}/downloads/${jar}/


# ADD GO DOCS

## start.sh
#!/bin/sh

java -Xms1G -Xmx1G -XX:+UseG1GC -XX:G1HeapRegionSize=4M -XX:+UnlockExperimentalVMOptions -XX:+ParallelRefProcEnabled -XX:+AlwaysPreTouch -XX:MaxInlineLevel=15 -jar velocity*.jar

java -Xms${JAVA_MEMORY} -Xmx${JAVA_MEMORY} -XX:+UseG${JAVA_MEMORY}C -XX:G1HeapRegionSize=4M -XX:+UnlockExperimentalVMOptions -XX:+ParallelRefProcEnabled -XX:+AlwaysPreTouch -XX:MaxInlineLevel=15 -jar /velocity/velocity

chmod +x start.sh
./start.sh


## veloctiy.toml
```toml
[servers]
# Configure your servers here. Each key represents the server's name, and the value represents the IP address of the server to connect to.
lobby = "0.0.0.0:25565"
vanilla = "0.0.0.0:25565"
modded = "0.0.0.0:25565"

# What order we should try to connect players to servers
try = [
    "lobby"
]
```

```toml
# Config version. Do not change this
config-version = "1.0"
# What port should the proxy be bound to? By default, we'll bind to all addresses on port 25577.
bind = "0.0.0.0:25577"
# What should be the MOTD? This gets displayed when the player adds your server to
# their server list. Legacy color codes and JSON are accepted.
motd = "&#09add3A Velocity Server"
# What should we display for the maximum number of players? (Velocity does not support a cap
# on the number of players online.)
show-max-players = 500
# Should we authenticate players with Mojang? By default, this is on.
online-mode = true
# If client's ISP/AS sent from this proxy is different from the one from Mojang's
# authentication server, the player is kicked. This disallows some VPN and proxy
# connections but is a weak form of protection.
prevent-client-proxy-connections = false
# Should we forward IP addresses and other data to backend servers?
# Available options:
# - "none":        No forwarding will be done. All players will appear to be connecting
#                  from the proxy and will have offline-mode UUIDs.
# - "legacy":      Forward player IPs and UUIDs in a BungeeCord-compatible format. Use this
#                  if you run servers using Minecraft 1.12 or lower.
# - "bungeeguard": Forward player IPs and UUIDs in a format supported by the BungeeGuard
#                  plugin. Use this if you run servers using Minecraft 1.12 or lower, and are
#                  unable to implement network level firewalling (on a shared host).
# - "modern":      Forward player IPs and UUIDs as part of the login process using
#                  Velocity's native forwarding. Only applicable for Minecraft 1.13 or higher.
player-info-forwarding-mode = "NONE"
# If you are using modern or BungeeGuard IP forwarding, configure a unique secret here.
forwarding-secret = "nBMS1GIk67Aw"
# Announce whether or not your server supports Forge. If you run a modded server, we
# suggest turning this on.
# 
# If your network runs one modpack consistently, consider using ping-passthrough = "mods"
# instead for a nicer display in the server list.
announce-forge = false
# If enabled (default is false) and the proxy is in online mode, Velocity will kick
# any existing player who is online if a duplicate connection attempt is made.
kick-existing-players = false
# Should Velocity pass server list ping requests to a backend server?
# Available options:
# - "disabled":    No pass-through will be done. The velocity.toml and server-icon.png
#                  will determine the initial server list ping response.
# - "mods":        Passes only the mod list from your backend server into the response.
#                  The first server in your try list (or forced host) with a mod list will be
#                  used. If no backend servers can be contacted, Velocity won't display any
#                  mod information.
# - "description": Uses the description and mod list from the backend server. The first
#                  server in the try (or forced host) list that responds is used for the
#                  description and mod list.
# - "all":         Uses the backend server's response as the proxy response. The Velocity
#                  configuration is used if no servers could be contacted.
ping-passthrough = "DISABLED"

[servers]
        # Configure your servers here. Each key represents the server's name, and the value
        # represents the IP address of the server to connect to.
        lobby = "127.0.0.1:30066"
        minigames = "127.0.0.1:30068"
        # In what order we should try servers when a player logs in or is kicked from a server.
        try = ["lobby"]
        factions = "127.0.0.1:30067"

[forced-hosts]
        "minigames.example.com" = ["minigames"]
        # Configure your forced hosts here.
        "lobby.example.com" = ["lobby"]
        "factions.example.com" = ["factions"]

[advanced]
        # Specify a custom timeout for connection timeouts here. The default is five seconds.
        connection-timeout = 5000
        # Enables BungeeCord plugin messaging channel support on Velocity.
        bungee-plugin-message-channel = true
        # Specify a read timeout for connections here. The default is 30 seconds.
        read-timeout = 30000
        # Enables TCP fast open support on the proxy. Requires the proxy to run on Linux.
        tcp-fast-open = false
        # Shows ping requests to the proxy from clients.
        show-ping-requests = false
        # By default, Velocity will attempt to gracefully handle situations where the user unexpectedly
        # loses connection to the server without an explicit disconnect message by attempting to fall the
        # user back, except in the case of read timeouts. BungeeCord will disconnect the user instead. You
        # can disable this setting to use the BungeeCord behavior.
        failover-on-unexpected-server-disconnect = true
        # How much compression should be done (from 0-9). The default is -1, which uses the
        # default level of 6.
        compression-level = -1
        # Declares the proxy commands to 1.13+ clients.
        announce-proxy-commands = true
        # Enables compatibility with HAProxy's PROXY protocol. If you don't know what this is for, then
        # don't enable it.
        haproxy-protocol = false
        # Enables the logging of commands
        log-command-executions = false
        # How large a Minecraft packet has to be before we compress it. Setting this to zero will
        # compress all packets, and setting it to -1 will disable compression entirely.
        compression-threshold = 256
        # How fast (in milliseconds) are clients allowed to connect after the last connection? By
        # default, this is three seconds. Disable this by setting this to 0.
        login-ratelimit = 3000

[query]
        # If query is enabled, on what port should the query protocol listen on?
        port = 25577
        # Whether plugins should be shown in query response by default or not
        show-plugins = false
        # This is the map name that is reported to the query services.
        map = "Velocity"
        # Whether to enable responding to GameSpy 4 query responses or not.
        enabled = false

# Legacy color codes and JSON are accepted in all messages.
[messages]
        generic-connection-error = "&cAn internal error occurred in your connection."
        already-connected = "&cYou are already connected to this proxy!"
        online-mode-only = "&cThis server only accepts connections from online-mode clients.\n\n&7Did you change your username? Sign out of Minecraft, sign back in, and try again."
        # Prefix when the player is disconnected from a server.
        #   First argument '%s': the server name
        disconnect-prefix = "&cCan't connect to %s: "
        no-available-servers = "&cThere are no available servers."
        # Prefix when the player gets kicked from a server.
        #   First argument '%s': the server name
        kick-prefix = "&cKicked from %s: "
        moved-to-new-server-prefix = "&cThe server you were on kicked you: "
```
