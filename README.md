# Abandoned, Not Forgotten — Earth Base Restoration

**Team:** Orbital Orphans  
**Challenge:** NASA Space Apps Challenge — Abandoned Equipment  
**Engine:** Defold 1.14.0 (Lua) — HTML5 / WASM  
**Play it here:** https://samartha-magar.itch.io/nasa-hackathon-abandoned-not-forgotten-project

> The Moon is full of machines we left behind. They are not junk. They are history.  
> Rebuild Earth Base, fly to the Moon, find them, and bring their stories home.


## Play

**Demo:** https://samartha-magar.itch.io/nasa-hackathon-abandoned-not-forgotten-project 
**Presentation (7 slides):** xxxx

### Controls

| Action | Key / Input |
| --- | --- |
| Move | WASD / Arrow Keys |
| Interact / Talk / Open Storybook | E |
| Click Shop Button | Mouse / Touch |
| Close Dialogue | E / Click |

### How to Play

1. Start at Earth Base. Talk to Commander Hayes.
2. Collect Logs, earn Coins and Science.
3. Open the Shop to build and upgrade your base.
4. Go to the Launch Pad. Wait for refuel (18s) and launch to the Moon.
5. Walk to the 6 abandoned equipment sites. Press **E** to open the Storybook for each one.
6. Return to Earth in the rocket to finish your run.

### The 6 Abandoned Machines

All stories in-game use real 【entity-NASA¦canonical_name=NASA】 history:

1. **Apollo 11** — Eagle Lunar Module Descent Stage (1969)
2. **Apollo 12** — Surveyor 3 camera brought home (1969)
3. **Apollo 15 LRV** — Lunar Roving Vehicle (1971)
4. **Surveyor 3** — Robotic lander (1967)
5. **Lunokhod 1** — Soviet rover (1970)
6. **Luna 2** — First human-made object to reach the Moon (1959)

## Why This Matters

Every Apollo, Surveyor, Luna, and Lunokhod site is still on the Moon today.  
We built this game so players don't just read about them — they walk to them, open a field journal, and learn what each machine did and why we left it there.

## NASA Data Used

- NASA Apollo Program History — https://www.nasa.gov/history/apollo/
- 【entity-NASA¦canonical_name=NASA】 Science: Moon Missions — https://science.nasa.gov/moon/
- 【entity-NASA¦canonical_name=NASA】 Lunar Reconnaissance Orbiter (LRO) imagery / landing site data — https://science.nasa.gov/mission/lro/
- Smithsonian / 【entity-NASA¦canonical_name=NASA】 descriptions for Surveyor, Luna, and Lunokhod missions

All mission dates, locations, and outcomes in the Storybook were checked against NASA sources.

## Run It Locally

### In Defold
1. Open Defold 1.14.0
2. Open this project
3. Press **Project > Build** (or F5)

### As HTML5 (same as itch.io)
1. In Defold: **Project > Bundle > HTML5**
2. Open the bundle folder in Terminal:

cd Nasa_hackathon_project_abandoned_not_forgotten
python3 -m http.server 8000

Open in browser: http://localhost:8000

## Project Structure
Code
main/
  opening/        - Intro / Commander Hayes
  hud/            - Storybook HUD, dialogue, shop button
  player/         - Player movement + camera
  earth / moon    - Earth Base + Moon collections
  store/          - Shop UI
  build_controller/ - Build / interact logic
assets/           - Pixel art, atlases, sounds

Show less
AI Use Disclosure
We used AI assistance for:

Debugging Defold / Lua errors
Generating placeholder pixel art, later edited by us
Drafting storybook dialogue, then rewritten and fact-checked by us
All game logic, level design, NASA fact-checking, and final submission decisions were done by the team.

Team — Orbital Orphans
Built in Kathmandu, Nepal for NASA Space Apps Challenge 2026.
