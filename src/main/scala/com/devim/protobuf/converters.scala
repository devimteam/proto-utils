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



  object timestamp {
    import com.google.protobuf.timestamp.Timestamp
    import java.time.{Instant,LocalDate, LocalDateTime, ZoneOffset}

    implicit class InstantConverter(val orig: Instant) extends AnyVal {
      def toProto: Timestamp = Timestamp(orig.getEpochSecond, orig.getNano)
    }

    implicit class TimestampProtoConverter(val orig: Timestamp) extends AnyVal {
      def toInstant: Instant = Instant.ofEpochSecond(orig.seconds, orig.nanos)
    }


    implicit class LocalDateConverter(val orig: LocalDate) extends AnyVal {
      def toProto(implicit zo: ZoneOffset = ZoneOffset.UTC): Timestamp =
        orig.atStartOfDay().toInstant(zo).toProto
    }

    implicit class LocalDateProtoConverter(val orig: Timestamp) extends AnyVal {
      def toLocalDate(implicit zo: ZoneOffset = ZoneOffset.UTC): LocalDate =
        LocalDateTime.ofInstant(orig.toInstant, zo).toLocalDate
    }
  }


}
