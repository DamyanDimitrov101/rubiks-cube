import React from 'react';

const CubeOperations = ({ handleReset, loading }) => {
    return (
        <div className="operation-buttons">
            <button
                onClick={handleReset}
                disabled={loading}
                className="operation-btn reset-btn"
            >
                Reset Cube
            </button>
        </div>
    );
};

export default CubeOperations;