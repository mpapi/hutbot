# Hutbot

Hutbot is a simple IRC bot that runs executable files (such as shell scripts)
to respond to commands. You can extend it by adding new scripts, and there's no
need to restart hutbot.


## Running

Assuming you have `$GOPATH` set,

    go get github.com/hut8labs/hutbot
    $GOPATH/bin/hutbot irc.example.com:6697 '#mychannel'


## Using

When hutbot joins your channel, you can talk to it:

    hutbot [~hutbot@example.com] entered the room.
    [you] hutbot: build myproject
    [hutbot] ok, building myproject

When you address hutbot, the first word after its name is interpreted as the
name of an executable to run (in hutbot's working directory) and the rest of
the line is interpreted as a single string argument.

So, the above message causes hutbot to run `./build`. The executable is called
with "myproject" on stdin, and with a bunch of environment variables set:

  * `$HUTBOT_SENDER` as the name of the sender ("you")
  * `$HUTBOT_CHANNEL` as the name of the channel ("#mychannel")
  * `$HUTBOT_CREATED` as the timestamp (seconds since epoch) for when the
    message was entered in IRC
  * `$HUTBOT_BOT` as the name of the bot ("hutbot")
  * `$HUTBOT_DIR` as hutbot's working directory
  * `$HUTBOT_COMMAND` as the first word of the message to hutbot ("build")
  * `$HUTBOT_ARGS` as the rest of the message to hutbot ("myproject")
  * `$HUTBOT_MESSAGE` as the full message ("hutbot: build myproject")

Whatever `./build` writes to stdout (in this case, "ok, building myproject") is
sent back to IRC. So, `./build` might contain:

    #!/bin/bash

    start-project-build ${HUTBOT_ARGS} && echo "ok, building ${HUTBOT_ARGS}" 

### Directories

If `./build` is a directory, hutbot will instead run each of the executable
files in that directory. (You'll see why you'd want this in a bit.)


### Special scripts

Hutbot also honors a few special script names.


#### IRC

These scripts let hutbot track who is in a channel.

  * `irc-join` is called when a user joins a channel
  * `irc-part` is called when a user parts a channel
  * `irc-quit` is called when a user quits a channel
  

#### Generic

These scripts allow hutbot to respond dynamically to messages.

  * `.all` is called for every message, even those not addressed to hutbot
    (except, for now, things that hutbot itself says)
  * `.missing` is called whenever there's no script for a command
   

#### Periodic

These scripts allow hutbot to take periodic actions (e.g., email a transcript
of a conversation after nobody has spoken for 30 minutes). They are invoked
with a subset of the usual environment variables (`$HUTBOT_BOT` and
`$HUTBOT_DIR`).

  * `.minute` is called once per minute
  * `.hour` is called once per hour
  * `.day` is called once per day
