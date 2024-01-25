N = 100
sink("housesOutputR.txt")
times <- system.time({
  for (i in 1:N) {
    houses = read.csv(file = "housesInput.csv", header = TRUE)
    print(summary(houses))
  }
})

sink()
print(times)
