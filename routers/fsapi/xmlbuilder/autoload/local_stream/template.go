package local_stream

var LOCAL_STREAM string = `<configuration name="local_stream.conf" description="stream files from local dir">
<!-- fallback to default if requested moh class isn't found -->
<directory name="default" path="$${sounds_dir}/music/8000">
  <param name="rate" value="8000"/>
  <param name="shuffle" value="true"/>
  <param name="channels" value="1"/>
  <param name="interval" value="20"/>
  <param name="timer-name" value="soft"/>
  <!-- list of short files to break in with every so often -->
  <!--<param name="chime-list" value="file1.wav,file2.wav"/>-->
  <!-- frequency of break-in (seconds)-->
  <!--<param name="chime-freq" value="30"/>-->
  <!-- limit to how many seconds the file will play -->
  <!--<param name="chime-max" value="500"/>-->
</directory>
</configuration>`
