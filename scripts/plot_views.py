import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv("assets/views.csv")
df["date"] = pd.to_datetime(df["timestamp"])
df = df.set_index("date")
df.plot(figsize=(8, 3), linewidth=2)
plt.savefig("assets/views.png", transparent=True, bbox_inches="tight", pad_inches=0.1)
plt.close()