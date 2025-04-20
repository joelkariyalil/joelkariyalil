import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv("assets/views.csv")
df["date"] = pd.to_datetime(df["timestamp"])
df = df.set_index("date")

ax = df.plot(
    figsize=(6, 3),
    linewidth=2.5,
    color="#00BFFF",
)

ax.tick_params(axis='both', which='major', colors='white')
ax.yaxis.label.set_color('white')
ax.xaxis.label.set_color('white')
ax.title.set_color('white')

for spine in ax.spines.values():
    spine.set_visible(False)

plt.savefig("assets/views.png", transparent=True, bbox_inches="tight", pad_inches=0.1)
plt.close()
