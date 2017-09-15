lazy val baseSettings = Seq(
  scalaVersion := "2.12.1",
  organization := "com.devim",
  version := "0.1.0-SNAPSHOT",
  resolvers ++= Seq(
    "Sonatype Nexus" at "https://nexus.devim.team/repository/maven-public/",
    "Central Proxy " at "https://nexus.devim.team/repository/maven-central/")
)

lazy val publishSettings = Seq(
  publishArtifact in Test := false,
  publishMavenStyle := true,
  publishTo := {
    val nexus = "https://nexus.devim.team/"
    if (isSnapshot.value)
      Some("snapshots" at nexus + "repository/maven-snapshots/")
    else
      Some("releases" at nexus + "repository/maven-releases/")
  },
  credentials += Credentials(Path.userHome / ".sbt" / ".credentials")
)


lazy val protoUtils = (project in file("."))
  .settings(baseSettings)
  .settings(publishSettings)
  .settings(
    name := "proto-utils",
    libraryDependencies ++= Seq(
      "com.google.protobuf" % "protobuf-java" % "3.3.1" %"protobuf",
      "org.scalatest" %% "scalatest" % "3.0.4" % "test"
    ),
    unmanagedResourceDirectories in Compile += baseDirectory.value /"src/main/protobuf",
    includeFilter in (Compile, unmanagedResources) := "*.proto",

    PB.protoSources in Compile := Seq(
      baseDirectory.value /"src/main/protobuf"),
    PB.targets in Compile := Seq(
      scalapb.gen() -> (sourceManaged in Compile).value
    )
  )

