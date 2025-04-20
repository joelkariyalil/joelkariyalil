import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv("assets/views.csv")
df["date"] = pd.to_datetime(df["timestamp"])
df = df.set_index("date")
df.plot()
plt.savefig("assets/views.png")