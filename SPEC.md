# Stars and Ships Specification

## 1. Project Summary
Stars and Ships is a web-based space strategy prototype where users navigate a 2D universe populated by planets and ships. The application runs in a web browser using Processing.js for rendering and interaction, served by a lightweight Go web server. It supports both desktop and mobile environments with touch-compatible controls. The core loop involves exploring the map, inspecting planets for resources, and observing autonomous ship movements.

## 2. User Interactions

### 2.1. Window / Canvas
- **Entry**: The game starts immediately upon loading the web page.
- **Splash Screen**: A startup image (`splashimg`) is displayed initially. Clicking it (`onClick`) calls `doStartup()` which hides the image and reveals the game canvas (`#sketch`).
- **Mobile optimization**: On mobile devices, the address bar is hidden via `window.scrollTo(0,1)`.

### 2.2. Mouse / Touch Controls
The following interactions are handled within `stars.pde`:

- **Left Click (Tap)**
    - **Select Planet**: If the cursor is within 15 pixels of a planet:
        - The planet index is selected.
        - A "Swirlie" animation appears around the planet.
        - The resource `HandleBox` (bar graph) becomes visible.
        - Console logs "planet [i] chosen".
    - **Select Ship**: If the cursor is within 15 pixels of a ship:
        - Toggles selection of that ship.
        - If selected, console logs "ship [i] chosen".
        - If deselected, console logs "ship [i] deselected".
    - **Move Target**: If no object is clicked:
        - Deselects current selection (`swirlie` and `handlebox` hidden).
        - Sets the global `PlanetX` and `PlanetY` (target coordinates) to the mouse position.
- **Right Click**
    - **Toggle View**: Toggles `overview` mode.
        - **Normal View**: Scale 2.0, Offset centered on mouse.
        - **Overview**: Scale 0.5, Offset (0,0).
- **Drag (Pan)**
    - **Map Movement**: Dragging the mouse updates the global offset (`offSetX`, `offSetY`) by the delta of mouse movement.
    - **Threshold**: Dragging is only registered if the total movement (`deltaX + deltaY`) exceeds 5 pixels.

### 2.3. Keyboard Controls
- **Debug Overlays**: While any key is pressed during the render loop:
    - Draws movement arrows (`drawArrow`) for all ships.
    - Draws halo rings (`drawHalo`) around ship destinations.
    - **Note**: The code uses `if (keyPressed)` inside the ship loop.

## 3. Major Datatypes

### 3.1. Objects (`Objects.pde`)

#### `Planet`
Encapsulates data for a single planet.
- **Fields**:
    - `x`, `y` (Float): Coordinates in the game world.
    - `owner` (Int): Index into the `player` array (0-3).
    - `name` (String): Name of the planet.
    - `resources` (Float[]): Array of 5 float values representing resources.
- **Constructor**: `Planet(float x, float y, int owner, String name, float[] resources)`

#### `Ship`
Encapsulates data for a single ship.
- **Fields**:
    - `name` (String): Name of the ship.
    - `speed` (Int): Speed of the ship.
    - `owner` (Int): Owner ID.
    - `x`, `y` (Float): Current position.
    - `destX`, `destY` (Float): Destination coordinates.
- **Constructor**: `Ship(String name, int speed, int owner, float x, float y, float destX, float destY)`

### 3.2. Global Arrays

#### `planets`
- **Data Structure**: Array of `Planet` objects.
- **Size**: 10.

#### `ships`
- **Data Structure**: Array of `Ship` objects.
- **Size**: 10.

#### `player`
Defines the factions in the game.
- **Data Structure**: Array of Arrays.
- **Format**: `[PlayerName, Color]`
    - `Color` is a hex integer (e.g., `#FF0000`).

### 3.3. Helper Objects (`stars.pde`)

#### `Swirlie`
A visual selector effect consisting of a trail of particles circling a target.
- **Fields**:
    - `num` (Int): Number of particles (20).
    - `mx[]`, `my[]` (Float Arrays): History of particle positions.
    - `x`, `y` (Int): Center position.
    - `visible` (Boolean): Render state.
    - `theta` (Float): Current angle.
- **Public Methods**:
    - `moveTo(newX, newY, newSelected)`: Updates position and selection state.
    - `redraw()`: Updates particle physics (rotation) and renders the trail.

#### `HandleBox`
A UI element displaying resource bars for a selected planet.
- **Fields**:
    - `handles[]` (Array of `Handle` objects).
    - `visible` (Boolean): Render state.
- **Public Methods**:
    - `setPos(newx, newy)`: Updates screen position relative to the selected object.
    - `redraw()`: Calls update/display on all child handles.
    - `setBar(index, newLength)`: Updates the value of a specific resource bar.

#### `Handle`
A single bar within the `HandleBox`.
- **Fields**:
    - `length` (Int): The value/length of the bar.
    - `description` (String): Label (e.g., "Mines", "Cows").
- **Public Methods**:
    - `update()`: Recalculates bounding box based on position.
    - `display()`: Draws the bar line and text label.

## 4. Algorithms

### 4.1. Initialization (`setup` / `makePlanets` / `makeShips`)
Pseudocode:
```
Set canvas size and framerate
Initialize player array
FOR i = 0 TO 9:
    res = [5 random floats]
    planets[i] = NEW Planet(randomX, randomY, randomOwner, nextPlanetName(), res)
    ships[i] = NEW Ship(nextShipName(), speed, randomOwner, randomX, randomY, randomDestX, randomDestY)
END FOR
Initialize HandleBox
```

### 4.2. Main Game Loop (`draw`)
Pseudocode:
```
Clear Screen
Apply View Scaling (viewscale)

// Handle "Camera" Ship
IF Distance(CameraX, CameraY, TargetX, TargetY) < 30:
    Select Random Planet P from planets[]
    TargetX = P.x
    TargetY = P.y

// Update Ships
FOR EACH ship in ships:
    IF Distance(ship.x, ship.y, ship.destX, ship.destY) < 30:
        Select Random Planet P from planets[]
        ship.destX = P.x
        ship.destY = P.y
    
    // Linear interpolation movement
    dx = ship.x - ship.destX
    dy = ship.y - ship.destY
    ratio = Min(Abs(dx)/Abs(dy), 10.0) / 10.0
    
    // Move ship towards destination
    ship.x += -ratio * Direction(dx)
    ship.y += -(inverse_ratio) * Direction(dy)
    
    DrawShip(ship)
END FOR

Draw Swirlie (if visible)
Draw Planets
Draw Resource Bars (if visible)
```

## 5. File Formats & External Files

### 5.1. Images
The application expects the following images in the same directory:
- `startup.png`: Splash screen.
- `stars_icon.png`, `stars-icon-ipad.png`: iOS icons.
- `stars.css`: Stylesheet.

### 5.2. Libraries
- `processing.js`: Core Processing engine.
- `jquery.js`: DOM manipulation.
- `stars.js`: Helper functions.
- `datasets.js`: Data generation helpers (`nextPlanet`, `nextShip`).
- `Objects.pde`: Class definitions for `Planet` and `Ship`.
