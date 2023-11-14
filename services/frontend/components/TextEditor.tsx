import React, { useState } from "react";

function TextEditor({  ws, startText}) {
  const [text, setText] = useState(startText);
  return (
    <div>
      <h2>
        TextEditor
      </h2>
      <textarea
        onChange={(e) => {
          ws.send(JSON.stringify({
            "title": "Test document",
            "body": e.target.value,
          }));
        }}
        value={text}
      />
    </div>
  );
}

export default TextEditor;
