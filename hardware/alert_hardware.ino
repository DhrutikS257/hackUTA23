#include <rgb_lcd.h>
#include "SR04.h"
#define TRIG_PIN 6
#define ECHO_PIN 5
SR04 sr04 = SR04(ECHO_PIN,TRIG_PIN);
const int buzzer = 8; //buzzer to arduino pin 8
const int pir = 6; //connected to arduino pin 6
const int vib = 2;
long a;
rgb_lcd lcd;

int alarmCount = 0;
int motionValue = 0;
int pirState = LOW;
int vibState = LOW;
const int colorR = 255;
const int colorG = 0;
const int colorB = 0;

typedef enum {
  warning, alert2, threat} AlertType_t;


void setup(){
 
  pinMode(buzzer, OUTPUT); // Set buzzer - pin 8 as an output
  pinMode(pir,INPUT); // Set buzzer - pin 6 as an input

  //set up LCD
  lcd.begin(16,2);
  lcd.setRGB(colorR,colorG,colorB);

  //set up serial Monitor
  Serial.begin(9600);


}
int UltraSonic()
{
   a = sr04.Distance();
   if(a < 150)
   {
     return 1;
   }
   delay(1000);
}
int VibrationSensor()
{
  int vibration = digitalRead(vib);
  if(vibration == HIGH)
  {
    return 1;
  }
  else 
  {
    return 0;
  }

}
int motionSensor()
{
  motionValue = digitalRead(pir);
  if(motionValue == HIGH)
  {
    if(pirState == LOW)
    {
      return 1;
    }
  }
  else
  {
    return 0;
  }
} 

void imposter(){
tone(buzzer, 1046); 
  delay(250);
  tone(buzzer, 1244); 
  delay(250); 
  tone(buzzer, 1400); 
  delay(250); 
  tone(buzzer, 1510); 
  delay(250);
  tone(buzzer, 1400); 
  delay(250);
  tone(buzzer, 1244); 
  delay(250); 
  tone(buzzer, 1046); 
  delay(250); 
  noTone(buzzer); 
  delay(500); 
  tone(buzzer, 932); 
  delay(125);
  tone(buzzer, 1174); 
  delay(125); 
  tone(buzzer, 1046); 
  delay(250);
  // end of first
  noTone(buzzer); 
  delay(500); 
  tone(buzzer, 780); 
  delay(250); 
  tone(buzzer, 525); 
  delay(250); 
  noTone(buzzer); 
  delay(250);
  //secont part
  tone(buzzer, 1046); 
  delay(250);
  tone(buzzer, 1244); 
  delay(250); 
  tone(buzzer, 1400); 
  delay(250); 
  tone(buzzer, 1510); 
  delay(250);
  tone(buzzer, 1400); 
  delay(250);
  tone(buzzer, 1244); 
  delay(250);
  tone(buzzer, 1400); 
  delay(250);
  noTone(buzzer); 
  delay(750);
  //fast part
  tone(buzzer, 1510); 
  delay(125);
  tone(buzzer, 1400); 
  delay(125);
  tone(buzzer, 1244); 
  delay(125);
  tone(buzzer, 1510); 
  delay(125);
  tone(buzzer, 1400); 
  delay(125);
  tone(buzzer, 1244); 
  delay(125);
  tone(buzzer, 1510); 
  delay(125);
  tone(buzzer, 1400); 
  delay(125);
  tone(buzzer, 1244); 
  delay(125);
  tone(buzzer, 1510); 
  delay(125);
  tone(buzzer, 1400); 
  delay(125);
  tone(buzzer, 1244); 
  delay(125);
   tone(buzzer, 1510); 
  delay(125);
  noTone(buzzer); 
  delay(500);
}
void LCDalert(AlertType_t alert){
  if(alert == threat)
  {
    lcd.clear();
    lcd.print("When the imposter is sus");
    delay(1000);
    lcd.clear();
  }
  if(alert == warning)
  {
    lcd.clear();
    lcd.print("Suspicious acti-");
    lcd.setCursor(1,10);
    lcd.print("vity detected");
    delay(1000);
    lcd.clear();

  }
  if(alert == alert2)
  {
    lcd.clear();
    lcd.print("Bro?");
    lcd.setCursor(1,10);
    lcd.print("Back up");
    delay(1000);
    lcd.clear();
  }
  if(alert == threat)
  {
    lcd.clear();
    lcd.print("GET OUT");
    lcd.setCursor(1,10);
    lcd.print("THE HOUSE");
    delay(1000);
    lcd.clear();
  }

}
void sensors(){
  //as soon as sensors detect breakin
  if(motionSensor()==1)
  {
    String data = "Motion sensor triggered";
    Serial.println(data);
    delay(500);
    LCDalert(warning); 
  }
  else if(UltraSonic() == 1)
  {
    String data = "Ultra Sonic triggered";
    Serial.println(data);
    delay(500);
    LCDalert(alert2); 
  }
  else if(VibrationSensor() == 1)
  {
    String data = "Vibration Sensor triggered";
    Serial.println(data);
    delay(500);
    LCDalert(threat);
  }

}

void loop(){
 
 //sensors running
 sensors();
 
}