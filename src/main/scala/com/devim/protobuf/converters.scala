package com.devim.protobuf

import scala.language.implicitConversions

object converters {
  object decimal {
    import com.devim.protobuf.math.{
    BigDecimal => ProtoDecimal,
    BigInteger => ProtoInt
    }
    import com.google.protobuf.ByteString

    implicit class BigDecimalConverter(val orig: BigDecimal) extends AnyVal {
      def toProto =
        ProtoDecimal(
          orig.scale,
          Some(ProtoInt(
            ByteString.copyFrom(orig.bigDecimal.unscaledValue().toByteArray))))
    }

    implicit class ProtoBigDecimalConverter(val orig: ProtoDecimal)
      extends AnyVal {
      def toBigDecimal =
        BigDecimal(orig.intValue
          .map(p => BigInt(p.value.toByteArray))
          .getOrElse(BigInt(0)),
          orig.scale)
    }

  }

}
