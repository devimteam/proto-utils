scalaVersion := "2.12.1"
organization := "com.devim"
version := "0.1.0-SNAPSHOT"

unmanagedResourceDirectories in Compile += baseDirectory.value / "src"
PB.protoSources in Compile := Seq(baseDirectory.value / "src")

PB.targets in Compile := Seq(
  scalapb.gen() -> (sourceManaged in Compile).value
)