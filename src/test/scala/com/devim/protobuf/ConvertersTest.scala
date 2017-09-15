package com.devim.protobuf

import org.scalatest.FlatSpec

class ConvertersTest extends FlatSpec {
  import converters.decimal._

  it should "serialise Scala BigDecimal to proto format and back" in {
    val initialData = BigDecimal.apply("10.01123124")

    val proto = initialData.toProto

    val resultData = proto.toBigDecimal

    assert(initialData===resultData)
  }
}
