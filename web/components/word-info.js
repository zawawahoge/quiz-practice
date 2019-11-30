import React from 'react'
import Link from 'next/link'

import Card from '@material-ui/core/Card';

import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';

import Box from '@material-ui/core/Box';
import { borders } from '@material-ui/system';


const WordInfo = (props) => (

  <Card onClick={props.onClick} style={{margin: '5%'}}>
    <CardHeader 
      title={props.data.word}   
      subheader={'https://ja.wikipedia.org/wiki/' + props.data.word}
    />
    <Box display={props.visible ? 'inline' : 'none'}>
      <CardContent>
        <div>
          {props.data.detail} 
        </div>
      </CardContent>
      
    </Box>
  </Card>
)

export default WordInfo
