function [] = plotCircle(x0, y0, R, linewidth, color, txt);
t=0:pi/180:2*pi;
x=R*cos(t)+x0; 
y=R*sin(t)+y0; 
%plot(x,y,'k');
%linewidth = 1;
fill(x,y, color, 'LineWidth', linewidth);
hold on;
text(x0 , y0, txt,'FontSize',7,'HorizontalAlignment','center');

