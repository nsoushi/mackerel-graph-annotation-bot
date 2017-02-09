child_process = require 'child_process'

module.exports = (robot) ->
  robot.respond /note\s+(.+)$/, (msg) ->
    option = msg.match[1]

    child_process.exec "graph-annotation post #{option}", (error, stdout, stderr) ->
          if !error
            output = stdout+''
            msg.send output
          else
            output = stderr+''
            msg.send output

