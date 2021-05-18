package fifo

// var FIFOCONF string = `<configuration name="fifo.conf" description="FIFO Configuration">
// <settings>
//   <param name="delete-all-outbound-member-on-startup" value="false"/>
//   <!--<param name="odbc-dsn" value="dsn:user:pass"/>-->
// </settings>
// <fifos>
//   <fifo name="cool_fifo@$${domain}" importance="0">
//     <!--<member timeout="60" simo="1" lag="20">{member_wait=nowait}user/1005@$${domain}</member>-->
//   </fifo>
// </fifos>
// </configuration>`

const FIFO string = `    <fifo name="%s" importance="%s">
      %s
    </fifo>`
const FIFOMEMBER string = `      <member timeout="%s" simo="%s" lag="%s">{member_wait=nowait}%s</member>
`
