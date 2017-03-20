scalaVersion := "2.12.1"
organization := "com.devim"
version := "0.1.0-SNAPSHOT"

unmanagedResourceDirectories in Compile += baseDirectory.value / "src"
PB.protoSources in Compile := Seq(baseDirectory.value / "src")

publishArtifact in Test := false
publishMavenStyle := true
publishTo := {
	val nexus = "http://nexus.devim.team/"
      	if (isSnapshot.value)
        	Some("snapshots" at nexus + "repository/maven-snapshots/")
	else
        	Some("releases" at nexus + "repository/maven-releases/")
}
credentials += Credentials(Path.userHome / ".sbt" / ".credentials")
resolvers ++= Seq(
      "Sonatype Nexus" at "http://nexus.devim.team/repository/maven-public/",
      "Central Proxy " at "http://nexus.devim.team/repository/maven-central/")

