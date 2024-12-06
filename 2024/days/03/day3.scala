import scala.util.matching.Regex

@main
def day03(fileName: String, pt2: Boolean): Unit =
  println(s"Starting day 2 with file: $fileName")

  val path: os.Path = os.pwd / fileName
  val lines: Seq[String] = os.read.lines(path)

  var sum = 0
  var isEnabled = true
  for l <- lines do
    val (v, e) = processLine(l, pt2, isEnabled)
    sum = sum + v
    isEnabled = e

  println(s"Sum: $sum")

def processLine(
    line: String,
    pt2: Boolean,
    enabled: Boolean
): (Int, Boolean) =
  var pattern = "(mul\\(\\d{1,3},\\d{1,3}\\))".r
  if pt2 then
    pattern = "(mul\\(\\d{1,3},\\d{1,3}\\))|(do\\(\\))|(don't\\(\\))".r
  val digitPattern = "\\d+".r
  val splits = pattern.findAllMatchIn(line).toList
  println(s"Splits: $splits")
  var sum = 0
  var isEnabled = enabled
  splits.foreach(s =>
    if s.toString() == "do()" then isEnabled = true
    else if s.toString() == "don't()" then isEnabled = false
    else
      // println(s"Processing digits $isEnabled")
      print("")
      if isEnabled then
        val digits = digitPattern
          .findAllMatchIn(s.toString())
          .map(d => d.toString().toInt)
          .toList
        // println(s"Digits: $digits")
        sum = sum + (digits(0) * digits(1))
  )

  return (sum, isEnabled)
