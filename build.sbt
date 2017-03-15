enablePlugins(ProtoSettings)
unmanagedResourceDirectories in Compile += baseDirectory.value / "src"
PB.protoSources in Compile := Seq(baseDirectory.value / "src")

 PB.targets in Compile := Seq(
      scalapb.gen(flatPackage=true) -> (sourceManaged in Compile).value
    )
