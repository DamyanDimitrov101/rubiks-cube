const API_URL = 'http://localhost:8080';

export const getCubeState = async () => {
    const response = await fetch(`${API_URL}/api/cube`);
    if (!response.ok) {
        throw new Error(`Failed to fetch cube state: ${response.statusText}`);
    }
    return response.json();
};

export const rotateFace = async (face, clockwise) => {
    const response = await fetch(`${API_URL}/api/cube/rotate`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ face, clockwise }),
    });
    if (!response.ok) {
        throw new Error(`Failed to rotate face: ${response.statusText}`);
    }
    return response.json();
};

export const resetCube = async () => {
    const response = await fetch(`${API_URL}/api/cube/reset`, {
        method: 'POST',
    });
    if (!response.ok) {
        throw new Error(`Failed to reset cube: ${response.statusText}`);
    }
    return response.json();
};
