import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv("assets/views.csv")
df["date"] = pd.to_datetime(df["timestamp"])
df = df.set_index("date")

# Plot with blue line and white text
ax = df.plot(
    figsize=(8, 3),
    linewidth=2.5,
    color="#00BFFF",  # Bright blue (you can tweak the hex)
)

# Make all text white
ax.tick_params(axis='both', which='major', colors='white')
ax.yaxis.label.set_color('white')
ax.xaxis.label.set_color('white')
ax.title.set_color('white')

# Hide borders
for spine in ax.spines.values():
    spine.set_visible(False)

# Save to transparent image
plt.savefig("assets/views.png", transparent=True, bbox_inches="tight", pad_inches=0.1)
plt.close()
