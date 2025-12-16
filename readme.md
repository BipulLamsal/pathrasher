## Ray Tracer

Rays are fired from the camera, through each pixel, into the scene
• When a ray hits an object, you compute a single surface point
• From that point, the renderer sends shadow rays toward each light source
• Those shadow rays are just visibility tests:
– if the light is visible → add its contribution
– if something blocks it → the light contributes nothing

• No light energy is bouncing around the scene
• No surface-to-surface illumination
• No indirect lighting
• No color bleeding

• Reflection and refraction rays may exist, but:
– they are limited in depth
– they still only gather direct lighting at each hit

So when the book says:

“the ray is then split into multiple rays that each go to a different light source”

What it means is:
• the renderer spawns multiple shadow rays
• one per light
• purely to check illumination
• not because light is physically splitting
