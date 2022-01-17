import React, { useState } from "react";
import { useForm } from "../hooks/useForm";
import encryptionService from "../services/encryptionsService";

export const RSAEncrypt = ({ encType, updateHistory }) => {
    const [decrypt, checkDecrypt] = useState(false);

    const [rsaValues, handleRSAChange, resetRSA] = useForm({
        exp: 0,
        mod: 0,
        text: ""
    });

    const handleEntryAdd = async () => {
        const response = await encryptionService.addEntry(
            {
                exp: parseInt(rsaValues.exp),
                mod: parseInt(rsaValues.mod),
                text: rsaValues.text,
                encrypt: !decrypt
            },
            encType.route
        );
        updateHistory(response.data);
    };

    const handleCheckBoxChange = () => {
        checkDecrypt(!decrypt);
    };

    return (
        <div className="encryptionSubmit">
            {encType.name}:<button onClick={resetRSA}>Reset Input</button>
            Exponent:{" "}
            <input
                className="exp"
                name="exp"
                value={rsaValues.exp}
                onChange={handleRSAChange}
            />
            Modulus:{" "}
            <input
                className="mod"
                name="mod"
                value={rsaValues.mod}
                onChange={handleRSAChange}
            />
            Text:{" "}
            <input
                className="text"
                name="text"
                value={rsaValues.text}
                onChange={handleRSAChange}
            />
            Decryption?:{" "}
            <input
                className="encrypt"
                name="encrypt"
                type="checkbox"
                onChange={handleCheckBoxChange}
            />
            <button onClick={handleEntryAdd}>Add</button>
        </div>
    );
};
