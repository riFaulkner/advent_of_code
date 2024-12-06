//> using scala 3.5.2
//> using toolkit 0.6.0

import scala.util.matching.Regex
import scala.collection.mutable.ArrayBuffer
import scala.compiletime.ops.double

@main
def day02(fileName: String, pt2: Boolean): Unit =
  println(s"Starting day 2 with file: $fileName")

  val path: os.Path = os.pwd / fileName
  val lines: Seq[String] = os.read.lines(path)

  var sum = 0
  for l <- lines do if processLine(l, pt2) then sum = sum + 1
  println(s"Sum: $sum")

def processLine(line: String, pt2: Boolean): Boolean =
  val pattern: Regex = "\\d+".r
  val splits =
    pattern
      .findAllMatchIn(line)
      .map(i => i.toString().toInt)
      .toList

  var isValid = isReportValid(splits, pt2)
  if !isValid && pt2 then
    splits.indices.foreach(i =>
      if !isValid then
        // println("Processing sub failures")
        val sub = ArrayBuffer[Int]().addAll(splits)
        sub.remove(i)
        if isReportValid(sub.toList, false) then isValid = true
    )

  return isValid

def isReportValid(report: List[Int], pt2: Boolean): Boolean =
  var r = report
  // If the report is in decending order, reverse it for simplicity
  if r(0) - r(1) > 0 then r = report.reverse

  var isValid = true
  var redos = 1
  var last = r(0)
  // now all lists are in what should be assending order.
  for i <- 1 to r.length - 1 if isValid do
    val diff = r(i) - last
    last = r(i)
    if diff < 1 || diff > 3 then
      isValid = false
      if pt2 && redos > 0 then
        redos = 0
        isValid = true
        last = r(i - 1)
  // println(s"Processing Line $r is valid -- $isValid")

  return isValid
