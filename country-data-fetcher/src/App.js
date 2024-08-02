import React, { useState } from 'react';
import axios from 'axios';
import { countries } from './countries';
import './App.css';

const App = () => {
  const [selectedCountry, setSelectedCountry] = useState('');
  const [data, setData] = useState('');

  const handleCountryChange = (event) => {
    setSelectedCountry(event.target.value);
  };

  const fetchData = async () => {
    try {
      if (!selectedCountry) {
        setData('No country selected');
        return;
      }

      const response = await axios.get(`https://country-items-api-ebc8cb908cd2.herokuapp.com/countries/${selectedCountry}`);
      setData(JSON.stringify(response.data, null, 2));
    } catch (error) {
      console.error('Error fetching data:', error);
      setData('Error fetching data');
    }
  };

  return (
    <div className="App">
      <h1>Fetch Country Data</h1>
      <form>
        <label>
          Select Country:
          <select value={selectedCountry} onChange={handleCountryChange}>
            <option value="">--Select a Country--</option>
            {countries.map(country => (
              <option key={country} value={country}>{country}</option>
            ))}
          </select>
        </label>
      </form>
      <button onClick={fetchData}>Fetch Data</button>
      <textarea value={data} readOnly rows="10" cols="50"></textarea>
    </div>
  );
};

export default App;
