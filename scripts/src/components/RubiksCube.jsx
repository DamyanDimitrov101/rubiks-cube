import React, { useState, useEffect } from 'react';
import './RubiksCube.css';
import { getCubeState, rotateFace, resetCube } from '../services/api';
import CubeNet from './CubeNet/CubeNet';
import FaceControls from './Controls/FaceControls';
import CubeOperations from './Controls/CubeOperations';

const RubiksCube = () => {
    const [cubeState, setCubeState] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(async() => {
        await fetchCubeState();
    }, []);

    const fetchCubeState = async () => {
        try {
            setLoading(true);
            const data = await getCubeState();
            setCubeState(data);
            setError(null);
        } catch (err) {
            console.error('Error fetching cube state:', err);
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    const handleRotate = async (face, clockwise) => {
        try {
            setLoading(true);
            const response = await rotateFace(face, clockwise);
            setCubeState(response.cube);
        } catch (error) {
            console.error('Error rotating face:', error);
            setError(`Failed to rotate face: ${error.message}`);
        } finally {
            setLoading(false);
        }
    };

    const handleReset = async () => {
        try {
            setLoading(true);
            const response = await resetCube();
            setCubeState(response.cube);
        } catch (error) {
            console.error('Error resetting cube:', error);
            setError(`Failed to reset cube: ${error.message}`);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="rubiks-cube-container">
            <h1>Rubik's Cube Simulator</h1>

            {error && (
                <div className="error-message">
                    <p>{error}</p>
                </div>
            )}

            {loading && <div className="loading-indicator">Loading...</div>}

            {cubeState && <CubeNet cubeState={cubeState} />}

            <div className="controls-container">
                <div className="control-section">
                    <h3>Rotate Faces</h3>
                    <div className="face-buttons row">
                        <FaceControls
                            handleRotate={handleRotate}
                            loading={loading}
                        />
                    </div>
                </div>

                <div className="control-section">
                    <h3>Cube Operations</h3>
                    <CubeOperations
                        handleReset={handleReset}
                        loading={loading}
                    />
                </div>
            </div>
        </div>
    );
};

export default RubiksCube;
