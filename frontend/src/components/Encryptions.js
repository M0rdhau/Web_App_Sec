import React, { useState, useEffect, useCallback } from "react";
import encryptionService from "../services/encryptionsService";
import { VigenereCaesar } from "./VigenereCaesar";
import { Response } from "./Response";
import { DiffieHellman } from "./DiffieHellman";
import { RSAEntry } from './RSAEntry';
import { RSAEncrypt } from './RSAEncrypt';

export const Encryptions = ({ encType }) => {
    const [history, setHistory] = useState([]);

    const getHistory = useCallback(async () => {
        const historyToSet = await encryptionService.getEntries(encType.route);
        setHistory(historyToSet.data);
    }, [encType]);

    const updateHistory = (newEntry) => {
        setHistory([newEntry].concat(history));
    };

    const deleteEntry = async (id) => {
        await encryptionService.deleteEntry(id, encType.route);
        setHistory(history.filter((each) => each.id !== id));
    };

    useEffect(() => {
        getHistory();
    }, [getHistory]);

    return (
        <div className="encryptionBox">
            <div>
                {(encType.name === "Caesar" || encType.name === "Vigenere") && (
                    <VigenereCaesar
                        encType={encType}
                        updateHistory={updateHistory}
                    />
                )}
                {encType.name === "Diffie-Hellmann" && (
                    <DiffieHellman
                        encType={encType}
                        updateHistory={updateHistory}
                    />
                )}
                {encType.name === "RSA - Key Generation" && (
                    <RSAEntry encType={encType} updateHistory={updateHistory} />
                )}
                {encType.name === "RSA - Encrypt/Decrypt" && (
                    <RSAEncrypt encType={encType} updateHistory={updateHistory} />
                )}
            </div>
            <div>
                {history.map((entry) => (
                    <Response
                        lastResponse={entry}
                        encType={encType}
                        deleteEntry={deleteEntry}
                    />
                ))}
            </div>
        </div>
    );
};
