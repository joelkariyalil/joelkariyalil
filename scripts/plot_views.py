import matplotlib.pyplot as plt
import matplotlib.dates as mdates
import pandas as pd

# Load and prepare data
df = pd.read_csv("assets/views.csv")
df["date"] = pd.to_datetime(df["timestamp"])
df = df.set_index("date")

# Plot
ax = df.plot(
    figsize=(6, 3),
    linewidth=2.5,
    color="#00BFFF",
)

# Format x-axis with month and year
ax.xaxis.set_major_formatter(mdates.DateFormatter('%b %Y'))
plt.setp(ax.get_xticklabels(), rotation=45, ha="right")

# Aesthetic settings
ax.tick_params(axis='both', which='major', colors='white')
ax.yaxis.label.set_color('white')
ax.xaxis.label.set_color('white')
ax.title.set_color('white')

for spine in ax.spines.values():
    spine.set_visible(False)

# Save the plot
plt.savefig("assets/views.png", transparent=True, bbox_inches="tight", pad_inches=0.1)
plt.close()
