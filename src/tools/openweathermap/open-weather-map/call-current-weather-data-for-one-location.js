/**
 * Function to get current weather data for a specified location.
 *
 * @param {Object} args - Arguments for the weather data request.
 * @param {string} [args.q] - City name for the weather query.
 * @param {number} [args.id] - City ID for the weather query.
 * @param {number} [args.lat] - Latitude of the location.
 * @param {number} [args.lon] - Longitude of the location.
 * @param {string} [args.zip] - Zip code for the weather query.
 * @param {string} [args.units] - Units for temperature (e.g., 'metric', 'imperial').
 * @param {string} [args.lang] - Language for the response.
 * @param {string} [args.Mode] - Format of the response (e.g., 'xml', 'html').
 * @param {string} [args.appid] - API key for authentication.
 * @returns {Promise<Object>} - The current weather data for the specified location.
 */
const executeFunction = async ({ q, id, lat, lon, zip, units, lang, Mode, appid }) => {
  const baseUrl = 'http://api.openweathermap.org/data/2.5';
  const apiKey = process.env.OPENWEATHERMAP_API_KEY;
  try {
    // Construct the URL with query parameters
    const url = new URL(`${baseUrl}/weather`);
    if (q) url.searchParams.append('q', q);
    if (id) url.searchParams.append('id', id);
    if (lat) url.searchParams.append('lat', lat);
    if (lon) url.searchParams.append('lon', lon);
    if (zip) url.searchParams.append('zip', zip);
    if (units) url.searchParams.append('units', units);
    if (lang) url.searchParams.append('lang', lang);
    if (Mode) url.searchParams.append('Mode', Mode);
    if (appid || apiKey) url.searchParams.append('appid', appid || apiKey);

    // Perform the fetch request
    const response = await fetch(url.toString(), {
      method: 'GET'
    });

    // Check if the response was successful
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData);
    }

    // Parse and return the response data
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error fetching current weather data:', error);
    return { error: 'An error occurred while fetching current weather data.' };
  }
};

/**
 * Tool configuration for fetching current weather data.
 * @type {Object}
 */
const apiTool = {
  function: executeFunction,
  definition: {
    type: 'function',
    function: {
      name: 'get_current_weather',
      description: 'Fetch current weather data for a specified location.',
      parameters: {
        type: 'object',
        properties: {
          q: {
            type: 'string',
            description: 'City name for the weather query.'
          },
          id: {
            type: 'integer',
            description: 'City ID for the weather query.'
          },
          lat: {
            type: 'number',
            description: 'Latitude of the location.'
          },
          lon: {
            type: 'number',
            description: 'Longitude of the location.'
          },
          zip: {
            type: 'string',
            description: 'Zip code for the weather query.'
          },
          units: {
            type: 'string',
            description: 'Units for temperature (e.g., metric, imperial).'
          },
          lang: {
            type: 'string',
            description: 'Language for the response.'
          },
          Mode: {
            type: 'string',
            description: 'Format of the response (e.g., xml, html).'
          },
          appid: {
            type: 'string',
            description: 'API key for authentication.'
          }
        }
      }
    }
  }
};

export { apiTool };