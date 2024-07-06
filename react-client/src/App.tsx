import React, { useState } from 'react';
import './App.css';
import { GreetingGenerationRequest } from './proto/greeting_service_pb';
import { ImageServiceClient } from './proto/Greeting_serviceServiceClientPb';

function App() {
  const [imagePrompt, setImagePrompt] = useState('');
  const [messagePrompt, setMessagePrompt] = useState('');
  const [imagePath, setImagePath] = useState('');
  const [error, setError] = useState<string | null>(null);

  const client = new ImageServiceClient("http://localhost:8080");
  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault(); // Prevent the default form submission behavior
    console.log("Form submitted"); // Debugging log

    if (messagePrompt.length > 32) {
      setError("Message prompt should be 32 characters or less");
      return;
    }
    setImagePath("loading_gif.gif")
    const req = new GreetingGenerationRequest();
    req.setImagePrompt(imagePrompt);
    req.setMessagePrompt(messagePrompt);

    try {
      const resp = await client.generateGreeting(req, {});
      console.log("Response received: ", resp.getImagePath()); // Debugging log
      setImagePath(resp.getImagePath());
      console.log("Setting image path to: ", resp.getImagePath())
    } catch (err) {
      console.log("somethings wrong")
      setError("Failed to generate greeting");
      console.error(err);
    }
    await new Promise(f => setTimeout(f, 5000));
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Greeting Generator</h1>
        <form onSubmit={handleSubmit}>
          <div className="input-container">
            <input
              type="text"
              placeholder="Enter image prompt"
              value={imagePrompt}
              onChange={(e) => setImagePrompt(e.target.value)}
            />
            <input
              type="text"
              placeholder="Enter message prompt (32 chars max)"
              value={messagePrompt}
              onChange={(e) => setMessagePrompt(e.target.value)}
              maxLength={32}
            />
            <button type="submit">Generate Greeting</button>
          </div>
        </form>
        {error && <p className="error">{error}</p>}
        {
          <div className="response">
            <img src={imagePath} alt="boop" />
          </div>
        }
      </header>
    </div>
  );
}

export default App;
