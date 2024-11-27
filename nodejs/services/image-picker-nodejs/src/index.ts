import express, { Request, Response } from 'express';
import { BUCKET_NAME } from './config';
import fs from 'fs';
import path from 'path';

let images : string[] = [];
try {
    const data = fs.readFileSync(path.resolve("images.json"), 'utf-8');
    const jsonData = JSON.parse(data);
    images = jsonData.images;
    images = images.map((filename) => `https://${BUCKET_NAME}.s3.amazonaws.com/${filename}`);
    console.log("Images loaded successfully:", images);
} catch (err) {
    console.error("Failed to load or parse images.json:", err);
    process.exit(1); // Exit the application
}

const app = express();
const PORT = process.env.PORT || 10116;

// Middleware to parse JSON bodies
app.use(express.json());

app.get("/health", (req: Request, res: Response) => {
    res.send("OK");
});

app.get('/imageUrl', async (req: Request, res: Response) => {
    const imageUrl = choose(images);
    res.send({ imageUrl });
});

function choose<T>(array: T[]): T {
    const i = Math.floor(Math.random() * array.length);
    return array[i];
}

// Start the server
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
